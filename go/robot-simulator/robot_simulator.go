package robot

import (
	"fmt"
	"strconv"
)

const (
	N Dir = iota
	E
	S
	W
)

func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

func Left() {
	Step1Robot.Dir = (Step1Robot.Dir + 3) % 4
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	}
}

func pointToDir(x, y int) Dir {
	panic("")
}

func (d Dir) String() string {
	return strconv.Itoa(int(d))
}

type Action struct {
	Cmd Command
}

func StartRobot(command chan Command, action chan Action) {
	for cmd := range command {
		action <- Action{Cmd: cmd}
	}
	close(action)
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	for act := range action {
		switch act.Cmd {
		case 'R':
			robot.Dir = (robot.Dir + 1) % 4

		case 'L':
			robot.Dir = (robot.Dir + 3) % 4

		case 'A':
			next := robot.Pos
			switch robot.Dir {
			case N:
				next.Northing++
			case E:
				next.Easting++
			case S:
				next.Northing--
			case W:
				next.Easting--
			}

			if next.Easting >= extent.Min.Easting &&
				next.Easting <= extent.Max.Easting &&
				next.Northing >= extent.Min.Northing &&
				next.Northing <= extent.Max.Northing {

				robot.Pos = next
			}
		}
	}

	report <- robot
}

type Action3 struct {
	Name string
	Cmd  Command
}

const (
	cmdStop Command = 0
)

func StartRobot3(name, script string, action chan Action3, _ chan string) {
	if name == "" {
		return
	}

	for _, c := range script {
		action <- Action3{
			Name: name,
			Cmd:  Command(c),
		}
	}

	action <- Action3{
		Name: name,
		Cmd:  cmdStop,
	}
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	defer close(rep)

	positions := make(map[string]Step2Robot)
	occupied := make(map[Pos]string)

	// Initial validation
	for _, rb := range robots {
		if rb.Name == "" {
			log <- "empty robot name"
			return
		}
		if _, exists := positions[rb.Name]; exists {
			log <- "duplicate name: " + rb.Name
			return
		}

		// Inline bounds check
		if rb.Pos.Easting < extent.Min.Easting ||
			rb.Pos.Easting > extent.Max.Easting ||
			rb.Pos.Northing < extent.Min.Northing ||
			rb.Pos.Northing > extent.Max.Northing {
			log <- rb.Name + " outside room"
			return
		}

		if _, taken := occupied[rb.Pos]; taken {
			log <- fmt.Sprintf("already occupied: %v", rb.Pos)
			return
		}

		positions[rb.Name] = rb.Step2Robot
		occupied[rb.Pos] = rb.Name
	}

	var result []Step3Robot

	for len(positions) > 0 {
		act := <-action

		r, ok := positions[act.Name]
		if !ok {
			log <- "unknown robot: " + act.Name
			return
		}

		switch act.Cmd {
		case 'A':
			next := r.Pos

			switch r.Dir {
			case N:
				next.Northing++
			case E:
				next.Easting++
			case S:
				next.Northing--
			case W:
				next.Easting--
			}

			// Wall check
			if next.Easting < extent.Min.Easting || next.Easting > extent.Max.Easting ||
				next.Northing < extent.Min.Northing || next.Northing > extent.Max.Northing {
				log <- "bumped into a wall at " + r.Dir.String()
				break
			}

			// Robot collision
			if other := occupied[next]; other != "" {
				log <- fmt.Sprintf("%v already occupied by %s", next, other)
				break
			}

			// Apply move
			delete(occupied, r.Pos)
			r.Pos = next
			occupied[r.Pos] = act.Name

		case 'R':
			r.Dir = (r.Dir + 1) % 4

		case 'L':
			r.Dir = (r.Dir + 3) % 4

		case cmdStop:
			delete(positions, act.Name)
			result = append(result, Step3Robot{act.Name, r})
			continue

		default:
			log <- "unknown action: " + string(act.Cmd)
			return
		}

		positions[act.Name] = r
	}

	rep <- result
}
