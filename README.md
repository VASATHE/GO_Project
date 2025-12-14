# Ticket Booking System (Go)

A **console-based Ticket Booking System** built using **Go (Golang)**.  
This project demonstrates core Go programming concepts such as **structs, slices, input validation, loops, and modular functions** by simulating a real-world ticket booking workflow.

---

## Project Overview
The Ticket Booking System allows users to book tickets by providing personal details.  
It validates user input at every step, updates ticket availability in real time, stores booking records, and displays all bookings once tickets are sold out.

This project was developed as a **learning-focused Go application** to strengthen backend fundamentals and control flow logic.

---

## Features
- Interactive console-based interface
- Step-by-step user input validation
  - First name & last name length validation
  - Email format validation
  - Ticket count validation
- Real-time ticket availability tracking
- Data storage using Go **structs and slices**
- Automatic program termination when tickets are sold out
- Displays complete booking summary at the end

---

## Tech Stack
- **Language:** Go (Golang)
- **Concepts Used:**
  - Structs
  - Slices
  - Loops & Conditionals
  - Functions
  - Constants & Variables
  - Input Validation

---

## Project Structure

    GO_Project/
    │
    ├── main.go
    └── README.md

---

## How to Run
1. Install Go from: https://go.dev/dl/
2. Clone the repository:
   ```bash
   git clone https://github.com/your-username/GO_Project.git
3. Navigate to the project directory:
   cd GO_Project
4. Run the program:
   go run main.go
   
---

## Sample Output

Welcome to Ticket Booking System

Total Tickets: 50
Currently available: 50

Enter your first name:
Vaishnavi
Enter your last name:
Sathe
Enter your email:
vaish@123
Enter the number of tickets:
40

Successfully booked tickets!
Remaining tickets are : 10

Enter your first name:
Priya
Enter your last name:
xyz
Enter your email:
pri@566
Enter the number of tickets:
10

Successfully booked tickets!
Remaining tickets are : 0

All tickets are booked!

List of booked tickets:
1) Vaishnavi Sathe vaish@123 - 40 tickets
2) Priya xyz pri@566 - 10 tickets