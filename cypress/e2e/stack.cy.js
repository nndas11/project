/// <reference types="cypress" />

describe('stack web interface', () => {
    beforeEach(() => {
        cy.visit('http://localhost:3000/')
      })

    it("check if page contains Create stack:",()=>{
        cy.contains("Create stack:")
    })

    it("Try to push item before stack declaration",()=>{
        cy.get('input[name="item"]').type('3')
        cy.get('#push-form').submit()
        cy.contains('Create stack')
    })

    it("Try to pop stack before declaration",()=>{
        cy.get('#pop-button').click()
        cy.contains("Create stack")
    })

    it("Try to get top item before stack declaration",()=>{
        cy.get('#top-button').click()
        cy.contains("Create stack")
    })

    it("Try to display stack before declaration",()=>{
        cy.get('#display-stack-button').click()
        cy.contains("Create stack")
    })

    it('check stack declaration', () => {
        cy.get('input[name="size"]').type('3')
        cy.get('#size-form').submit()
        cy.contains('Stack declared')

    })

    it('Try to display stack with no items',()=>{
        cy.get('#display-stack-button').click()
        cy.contains("Stack is empty:")
    })

    it('Try to get pop item when stack is empty',()=>{
        cy.get('#pop-button').click()
        cy.contains("Stack is empty")
    })

    it('Try to get top element when stack is empty',()=>{
        cy.get('#top-button').click()
        cy.contains('Stack is empty')
    })

    it("check element push to stack",()=>{
        cy.get('input[name="item"]').type('3')
        cy.get('#push-form').submit()
        cy.contains('Item added to stack')

    })

    it("check top element of stack",()=>{
        cy.get('#top-button').click()
        cy.contains('Top Item : 3')
    })

    it("Display Stack",()=>{
        cy.get('input[name="item"]').type('4')
        cy.get('#push-form').submit()
        cy.get('#display-stack-button').click()
        cy.contains("Stack Content: 4 3")
    })

    it("Pop Item from stack",()=>{
        cy.get('#pop-button').click()
        cy.contains("Popped item : 4")
    })
    
})