# CBLOL Schedule App

![CBLOL Logo](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ-Y4A29Fo3SfI8AeFJax3Uj9ulXumF388FWYtkf0QDBg&s)

## Overview

This application showcases the schedule for the Campeonato Brasileiro de League of Legends (CBLOL) in 2024.

**Note:** Bloated with unnecessary stuff just to improve Go with serverless learnings.

## Features

- **Games Schedule**: View the date, time, teams, and current wins x loses of each team for each CBLOL game in 2024.
- **Create Game Schedule**: Add scheduled games for upcoming weeks with their date and times.

## Technologies Used

- Go (Programming Language)
- [Go Templ](https://templ.guide/)
- [Echo](https://echo.labstack.com/)
- [Serverless Framework](https://www.serverless.com/)

## How to Run

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Davilopesm/cblol-schedule-go.git
   ```

2. **Navigate to the project directory:**

   ```bash
   cd cblol-schedule-go
   npm install
   ```

3. **Build and Run the application locally:**

   ```bash
   make && sls offline
   ```

4. **Open your web browser and visit [http://localhost:3000/local/games](http://localhost:3000/local/games) to view the CBLOL schedule.**

5. **Send a POST request to [http://localhost:3000/local/games](http://localhost:3000/local/games) to ADD games.**
   5.1 **An example request can be seem in model/insert_game_example.json**

6. **Build and Deploy to AWS**

   ```bash
   make deploy
   ```
