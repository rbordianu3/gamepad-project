package main

import (
    "context"
    "log"
    "time"

    "github.com/rbordianu3/gamepad-project/matrix/device"
    "github.com/richardlt/matrix/sdk-go/player"
)

func main() {
    api := &player.API{} // Assuming player.API is correctly initialized
    g := device.NewGamepad()

    if err := g.Init(api); err != nil {
        log.Fatalf("Failed to initialize gamepad: %v", err)
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    go func() {
        time.Sleep(10 * time.Second)
        cancel()
    }()

    if err := g.OpenDevices(ctx); err != nil {
        log.Fatalf("Failed to open devices: %v", err)
    }
}