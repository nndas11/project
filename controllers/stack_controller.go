package controller

import (
	"fmt"
	"proj/models"
	"strconv"

	"github.com/gofiber/fiber"
)

type StackController struct {
	stack *models.Stack
}

func NewStackController(size int) *StackController {
	return &StackController{
		stack: models.NewStack(size),
	}
}

// parse the json from request body
type RequestBody struct {
	Number string `json:"number"`
}

type Size struct {
	size int `json:size`
}

func (sc *StackController) Push(c *fiber.Ctx) {

	// item := new(Item)
	// err := c.Body()
	// number, err := strconv.Atoi(c.FormValue("item"))
	var requestBody RequestBody
	if err := c.BodyParser(&requestBody); err != nil {
		fmt.Println(requestBody.Number)
		c.Status(500).Send("Invalid item value")
		return
	}
	num, err := strconv.Atoi(requestBody.Number)
	fmt.Println(num)
	if err != nil {
		c.Status(fiber.StatusBadRequest).Send("Invalid item value")
		return
	}
	if sc.stack.Push(num) {
		c.Send("Item added to stack")
	} else {
		c.Status(fiber.StatusBadRequest).Send("Stack is full")
	}
}

func (sc *StackController) Pop(c *fiber.Ctx) {
	item, ok := sc.stack.Pop()
	fmt.Print(item)
	if ok {
		c.Send("Popped item: " + strconv.Itoa(item))
	} else {
		c.Status(fiber.StatusBadRequest).Send("Stack is empty")
	}
}

func (sc *StackController) Top(c *fiber.Ctx) {
	item, ok := sc.stack.Top()
	if ok {
		c.Send("Top item: " + strconv.Itoa(item))
	} else {
		c.Status(fiber.StatusBadRequest).Send("Stack is empty")
	}
}

func (sc *StackController) Display(c *fiber.Ctx) {
	if sc.stack.GetTop() == -1 {
		c.Status(fiber.StatusBadRequest).Send("Stack is empty")
		return
	}
	data := sc.stack.GetData()[:sc.stack.GetTop()+1] // get all elements up to the top of the stack
	c.JSON(data)
}
