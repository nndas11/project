package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
)

type mockStack struct {
	items []int
}

func (s *mockStack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *mockStack) Pop() int {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *mockStack) Top() int {
	if s.IsEmpty() {
		return 0
	}
	return s.items[len(s.items)-1]
}

func (s *mockStack) IsEmpty() bool {
	return len(s.items) == 0
}

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    int    `json:"data"`
}

func TestPopAPI(t *testing.T) {
	app := fiber.New()
	stack := &mockStack{}

	// stackController := controller.StackController{}
	app.Delete("/pop", func(c *fiber.Ctx) {
		if stack.IsEmpty() {
			c.Status(fiber.StatusOK).JSON(fiber.Map{
				"code":    fiber.StatusOK,
				"message": "Stack is empty.",
				"data":    "",
			})

		} else {
			c.Status(fiber.StatusOK).JSON(fiber.Map{
				"code":    fiber.StatusOK,
				"message": "Popped item",
				"data":    stack.Pop(),
			})
		}
	})
	t.Run("Pop item from stack", func(t *testing.T) {
		stack.Push(5)
		req := httptest.NewRequest("DELETE", "/pop", nil)
		res, err := app.Test(req)
		var rsp response
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &rsp)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, 5, rsp.Data)
		assert.Equal(t, "Popped item", rsp.Message)
	})

	t.Run("Pop item from empty stack", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/pop", nil)
		res, err := app.Test(req)
		var rsp response
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &rsp)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "Stack is empty.", rsp.Message)
	})
}

type request struct {
	Item int `json:"item"`
}

func TestPushAPI(t *testing.T) {
	app := fiber.New()
	stack := &mockStack{}

	app.Post("/push", func(c *fiber.Ctx) {
		var body request
		if err := c.BodyParser(&body); err != nil {
			c.JSON(fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": "Give number.",
				"data":    "",
			})
		}
		stack.Push(body.Item)
		c.JSON(fiber.Map{
			"code":    fiber.StatusCreated,
			"message": "Item added to stack.",
			"data":    "",
		})
	})

	t.Run("Push item to stack", func(t *testing.T) {
		reqBody := request{Item: 5}
		reqBodyBytes, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/push", bytes.NewBuffer(reqBodyBytes))
		req.Header.Add("Content-Type", "application/json")
		res, err := app.Test(req)
		assert.Nil(t, err)
		body, _ := ioutil.ReadAll(res.Body)
		var rsp response
		json.Unmarshal(body, &rsp)
		assert.Equal(t, http.StatusCreated, rsp.Code)
		assert.Equal(t, "Item added to stack.", rsp.Message)
	})

}

type displayResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []int  `json:"data"`
}

func TestDisplayAPI(t *testing.T) {
	app := fiber.New()
	stack := &mockStack{}

	// stackController := controller.StackController{}
	app.Get("/display", func(c *fiber.Ctx) {
		if stack.IsEmpty() {
			c.Status(fiber.StatusOK).JSON(fiber.Map{
				"code":    fiber.StatusOK,
				"message": "Stack is empty.",
				"data":    "",
			})

		} else {
			c.Status(fiber.StatusOK).JSON(fiber.Map{
				"code":    fiber.StatusOK,
				"message": "Stack contents",
				"data":    stack.items,
			})
		}
	})

	t.Run("Display empty stack", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/display", nil)
		res, err := app.Test(req)
		var rsp displayResponse
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &rsp)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "Stack is empty.", rsp.Message)
	})

	t.Run("Display content of stack", func(t *testing.T) {
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		req := httptest.NewRequest("GET", "/display", nil)
		res, err := app.Test(req)
		assert.Nil(t, err)
		body, _ := ioutil.ReadAll(res.Body)
		var rsp displayResponse
		json.Unmarshal(body, &rsp)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "Stack contents", rsp.Message)
		assert.Equal(t, []int{1, 2, 3}, rsp.Data)
	})
}

func TestTopAPI(t *testing.T) {
	app := fiber.New()
	stack := &mockStack{}

	app.Get("/top", func(c *fiber.Ctx) {
		if stack.IsEmpty() {
			c.JSON(fiber.Map{
				"code":    fiber.StatusOK,
				"message": "Stack is empty.",
				"data":    "",
			})
		} else {
			c.JSON(fiber.Map{
				"code":    fiber.StatusOK,
				"message": "Top element of stack",
				"data":    stack.Top(),
			})
		}
	})

	t.Run("Peek top element of empty stack", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/top", nil)
		res, err := app.Test(req)
		assert.Nil(t, err)
		body, _ := ioutil.ReadAll(res.Body)
		var rsp response
		json.Unmarshal(body, &rsp)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "Stack is empty.", rsp.Message)
	})

	t.Run("Peek top element of stack", func(t *testing.T) {
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		req := httptest.NewRequest("GET", "/top", nil)
		res, err := app.Test(req)
		assert.Nil(t, err)
		body, _ := ioutil.ReadAll(res.Body)
		var rsp response
		json.Unmarshal(body, &rsp)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "Top element of stack", rsp.Message)
		assert.Equal(t, 3, rsp.Data)
	})

}
