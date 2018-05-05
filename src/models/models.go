package main

import "time"

// User 
type User struct {
        Id       int        `json:"id"`
        Username string     `json:"username"`
        Email    string     `json:"email"`
        Phone    int `json:"phone"`
}

// order
type order struct {
        User        User       `json:"user"`
        id           int `json:"id"`
        orderNumber int `json:"orderNumber"`
        burgerType      string      `json:"burgerType"`
        yesBun    bool    `json:"bun"`
        proteinSize int `json:"proteinSize"`
        cheese     string `json:"cheese"`
        sauce string `json:"sauce"`
        topping string `json:"topping"`
        premiumtopping string `json:"premiumTopping"`
        bunType string `json:"bunType"`
        Timestamp time.Time `json:"timestamp"`
}

// process
type process struct {
        Id        int       `json:"id"`
        orderNumber      User      `json:"user"`
        Timestamp     time.Time    `json:"timestamp"`
        Comment   Comment   `json:"comment"`
        PaymentType string    `json:"paymentType"`
        Timestamp time.Time `json:"timestamp"`
}

type User    []User
type order []order
type process    []process