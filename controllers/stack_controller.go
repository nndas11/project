package controller

import (
	"fmt"
	"proj/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type StackController struct {
	stack *models.Stack
}

// func NewStackController(size int) *StackController {
// 	return &StackController{
// 		stack: models.NewStack(size),
// 	}
// }

// parse the json from request body
type RequestBody struct {
	Number string `json:"number"`
}

func (sc *StackController) Push(c *fiber.Ctx) error {

	//check whether stack is declared or not
	if sc.stack == nil {
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Create stack",
			"content": "",
		})
		return nil
	}

	var requestBody RequestBody
	if err := c.BodyParser(&requestBody); err != nil {
		fmt.Println(requestBody.Number)
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid item value",
			"content": "",
		})
		return nil
	}
	num, err := strconv.Atoi(requestBody.Number)
	fmt.Println(num)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid item value",
			"content": "",
		})
		return nil
	}
	if sc.stack.Push(num) {
		c.Status(fiber.StatusCreated).JSON(&fiber.Map{
			"code":    fiber.StatusCreated,
			"message": "Item added to stack",
			"content": "",
		})

	} else {
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Stack is full",
			"content": "",
		})

	}
	return nil
}

func (sc *StackController) Pop(c *fiber.Ctx) error {

	//check whether stack is declared or not
	if sc.stack == nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Create stack",
			"content": "",
		})

	}

	item, ok := sc.stack.Pop()
	if ok {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"code":    fiber.StatusOK,
			"message": "Popped item : ",
			"content": strconv.Itoa(item),
		})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Stack is empty",
			"content": "",
		})
	}
}

func (sc *StackController) Top(c *fiber.Ctx) error {

	//check whether stack is declared or not
	if sc.stack == nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Create stack",
			"content": "",
		})

	}

	item, ok := sc.stack.Top()
	if ok {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"code":    fiber.StatusOK,
			"message": "Top Item : ",
			"content": strconv.Itoa(item),
		})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"code":    fiber.StatusOK,
			"message": "Stack is empty",
			"content": "",
		})
	}
}

func (sc *StackController) Display(c *fiber.Ctx) error {

	//check whether stack is declared or not
	if sc.stack == nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Create stack",
			"content": "",
		})

	}

	data := sc.stack.Display()

	if data == nil {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"code":    fiber.StatusOK,
			"message": "Stack is empty",
			"content": "",
		})

	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Stack Content",
		"content": data,
	})
}

func (sc *StackController) Declare(c *fiber.Ctx) error {
	var requestBody RequestBody
	if err := c.BodyParser(&requestBody); err != nil {
		fmt.Println(requestBody.Number)
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid item value",
			"content": "",
		})

	}
	size, err := strconv.Atoi(requestBody.Number)
	fmt.Println(size)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid item value",
			"content": "",
		})

	}
	sc.stack = &models.Stack{}
	sc.stack.NewStack(size)

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "Stack declared",
		"content": "",
	})
}
