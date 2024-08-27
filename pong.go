package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameState int

const (
	StateHomeScreen GameState = iota
	StateCredits
	StateLevel1
)

const (
	screenWidth  = 800
	screenHeight = 450
	maxParticles = 100
)

type Particle struct {
	x, y   float32
	speedY float32
	color  rl.Color
}

var (
	currentState GameState = StateHomeScreen
	particles    [maxParticles]Particle
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Pong with Scenes")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	// Inicializar partículas
	for i := 0; i < maxParticles; i++ {
		particles[i] = Particle{
			x:      float32(rl.GetRandomValue(0, screenWidth)),
			y:      float32(rl.GetRandomValue(0, screenHeight)),
			speedY: float32(rl.GetRandomValue(50, 150)) / 100.0, // Convertir a float32
			color:  rl.DarkGray,                                 // Color gris oscuro para las partículas
		}
	}

	for !rl.WindowShouldClose() {
		switch currentState {
		case StateHomeScreen:
			homeScreen()
		case StateCredits:
			creditsScreen()
		case StateLevel1:
			level1()
		}
	}
}

func homeScreen() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	// Dibujar botones
	startBtn := rl.Rectangle{X: 350, Y: 150, Width: 100, Height: 50}
	creditsBtn := rl.Rectangle{X: 350, Y: 250, Width: 100, Height: 50}

	rl.DrawRectangleRec(startBtn, rl.LightGray)
	rl.DrawRectangleRec(creditsBtn, rl.LightGray)

	rl.DrawText("Comenzar", int32(startBtn.X+10), int32(startBtn.Y+15), 20, rl.DarkGray)
	rl.DrawText("Créditos", int32(creditsBtn.X+10), int32(creditsBtn.Y+15), 20, rl.DarkGray)

	// Chequear si se ha hecho clic en alguno de los botones
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		mousePoint := rl.GetMousePosition()

		if rl.CheckCollisionPointRec(mousePoint, startBtn) {
			currentState = StateLevel1
		} else if rl.CheckCollisionPointRec(mousePoint, creditsBtn) {
			currentState = StateCredits
		}
	}

	rl.EndDrawing()
}

func creditsScreen() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.DrawText("Créditos", 350, 50, 30, rl.DarkGray)
	rl.DrawText("Desarrollado por Agustin en agosto", 200, 200, 20, rl.DarkGray)
	rl.DrawText("Presiona ESC para volver", 250, 300, 20, rl.DarkGray)

	if rl.IsKeyPressed(rl.KeyEscape) {
		currentState = StateHomeScreen
	}

	rl.EndDrawing()
}

func level1() {
	// Variables del juego Pong
	paddleWidth := int32(20)
	paddleHeight := int32(100)
	player1Y := float32(screenHeight)/2 - float32(paddleHeight)/2
	player2Y := float32(screenHeight)/2 - float32(paddleHeight)/2

	ballX := float32(screenWidth / 2)
	ballY := float32(screenHeight / 2)
	ballRadius := float32(10.0)
	ballSpeedX := float32(300.0)
	ballSpeedY := float32(300.0)

	playerSpeed := float32(300.0)

	for !rl.WindowShouldClose() && currentState == StateLevel1 {
		// Update player 1 (left paddle) position
		if rl.IsKeyDown(rl.KeyW) && player1Y > 0 {
			player1Y -= playerSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyS) && player1Y < screenHeight-float32(paddleHeight) {
			player1Y += playerSpeed * rl.GetFrameTime()
		}

		// Update player 2 (right paddle) position
		if rl.IsKeyDown(rl.KeyUp) && player2Y > 0 {
			player2Y -= playerSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyDown) && player2Y < screenHeight-float32(paddleHeight) {
			player2Y += playerSpeed * rl.GetFrameTime()
		}

		// Update ball position
		ballX += ballSpeedX * rl.GetFrameTime()
		ballY += ballSpeedY * rl.GetFrameTime()

		// Ball collision with top and bottom
		if ballY-ballRadius <= 0 || ballY+ballRadius >= screenHeight {
			ballSpeedY *= -1
		}

		// Ball collision with paddles
		if (ballX-ballRadius <= float32(paddleWidth) && ballY >= player1Y && ballY <= player1Y+float32(paddleHeight)) ||
			(ballX+ballRadius >= screenWidth-float32(paddleWidth) && ballY >= player2Y && ballY <= player2Y+float32(paddleHeight)) {
			ballSpeedX *= -1
		}

		// Ball out of bounds
		if ballX-ballRadius <= 0 || ballX+ballRadius >= screenWidth {
			ballX = float32(screenWidth / 2)
			ballY = float32(screenHeight / 2)
			ballSpeedX = 300
			ballSpeedY = 300
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Dibujar partículas (nieve)
		for i := 0; i < maxParticles; i++ {
			particles[i].y += particles[i].speedY

			if particles[i].y > screenHeight {
				particles[i].y = 0
				particles[i].x = float32(rl.GetRandomValue(0, screenWidth))
				particles[i].speedY = float32(rl.GetRandomValue(50, 150)) / 100.0
			}

			rl.DrawCircle(int32(particles[i].x), int32(particles[i].y), 2, particles[i].color)
		}

		// Draw paddles
		rl.DrawRectangle(0, int32(player1Y), paddleWidth, paddleHeight, rl.Black)
		rl.DrawRectangle(screenWidth-paddleWidth, int32(player2Y), paddleWidth, paddleHeight, rl.Black)

		// Draw ball
		rl.DrawCircle(int32(ballX), int32(ballY), ballRadius, rl.Red)

		rl.EndDrawing()

		// Presiona ESC para volver a la pantalla de inicio
		if rl.IsKeyPressed(rl.KeyEscape) {
			currentState = StateHomeScreen
		}
	}
}
