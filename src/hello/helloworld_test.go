package main

import (
  "testing"
  "time"
)

func TestTime(t *testing.T) {
  current_time := time.Now()
  same_time := current_time
  if current_time != same_time {
    t.Errorf("It's very strange... It must be the same time: %v and %v",
      current_time, same_time)
  }
}

func TestTwo(t *testing.T) {
  if 2 != 2 {
    t.Errorf("It's not funny. Check your computer!")
  }
}
