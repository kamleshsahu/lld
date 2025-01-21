package main

import (
	"fmt"
	"time"
)

// Alarm represents an alarm with time, label, and a flag for whether it's active.
type Alarm struct {
	ID         int
	Time       time.Time
	Label      string
	IsActive   bool
	IsReminder bool // to distinguish between alarms and reminders
}

// AlarmClock holds the list of alarms and reminders.
type AlarmClock struct {
	alarms    []Alarm
	reminders []Alarm
}

// NewAlarmClock creates a new AlarmClock instance.
func NewAlarmClock() *AlarmClock {
	return &AlarmClock{}
}

// AddAlarm adds a new alarm.
func (ac *AlarmClock) AddAlarm(alarmTime time.Time, label string) int {
	alarm := Alarm{
		ID:         len(ac.alarms) + len(ac.reminders), // Simple ID generation
		Time:       alarmTime,
		Label:      label,
		IsActive:   true,
		IsReminder: false,
	}
	ac.alarms = append(ac.alarms, alarm)
	ac.scheduleAlarm(alarm)
	return alarm.ID
}

// AddReminder adds a new reminder.
func (ac *AlarmClock) AddReminder(reminderTime time.Time, label string) int {
	reminder := Alarm{
		ID:         len(ac.alarms) + len(ac.reminders), // Simple ID generation
		Time:       reminderTime,
		Label:      label,
		IsActive:   true,
		IsReminder: true,
	}
	ac.reminders = append(ac.reminders, reminder)
	ac.scheduleReminder(reminder)
	return reminder.ID
}

// DeleteAlarm removes an alarm by ID.
func (ac *AlarmClock) DeleteAlarm(id int) {
	for i, alarm := range ac.alarms {
		if alarm.ID == id {
			ac.alarms = append(ac.alarms[:i], ac.alarms[i+1:]...)
			fmt.Println("Alarm deleted:", alarm.Label)
			return
		}
	}
}

// ToggleAlarm toggles an alarm on or off by ID.
func (ac *AlarmClock) ToggleAlarm(id int, status bool) {
	for i, alarm := range ac.alarms {
		if alarm.ID == id {
			ac.alarms[i].IsActive = status
			fmt.Println("Alarm toggled:", alarm.Label, "Active:", status)
			return
		}
	}
}

// Schedule the alarm to ring after the specified time.
func (ac *AlarmClock) scheduleAlarm(alarm Alarm) {
	duration := alarm.Time.Sub(time.Now())
	if duration > 0 {
		time.AfterFunc(duration, func() {
			if alarm.IsActive {
				ac.ringAlarm(alarm)
			}
		})
	}
}

// Schedule the reminder to notify at the specified time.
func (ac *AlarmClock) scheduleReminder(reminder Alarm) {
	duration := reminder.Time.Sub(time.Now())
	if duration > 0 {
		time.AfterFunc(duration, func() {
			if reminder.IsActive {
				ac.ringReminder(reminder)
			}
		})
	}
}

// Ring the alarm when the scheduled time is reached.
func (ac *AlarmClock) ringAlarm(alarm Alarm) {
	fmt.Printf("ALARM RINGING: %s at %s\n", alarm.Label, alarm.Time.Format("15:04:05"))
}

// Ring the reminder when the scheduled time is reached.
func (ac *AlarmClock) ringReminder(reminder Alarm) {
	fmt.Printf("REMINDER: %s at %s\n", reminder.Label, reminder.Time.Format("15:04:05"))
}

// Distinguish between an alarm and a reminder.
func (ac *AlarmClock) DistinguishAlarmReminder(id int) string {
	for _, alarm := range ac.alarms {
		if alarm.ID == id {
			return "Alarm"
		}
	}
	for _, reminder := range ac.reminders {
		if reminder.ID == id {
			return "Reminder"
		}
	}
	return "Not Found"
}

func main() {
	ac := NewAlarmClock()

	// Add some alarms and reminders
	alarmID := ac.AddAlarm(time.Now().Add(5*time.Second), "Morning Alarm")
	reminderID := ac.AddReminder(time.Now().Add(10*time.Second), "Take Medicine")

	// Print out the types
	fmt.Println("Distinguishing Alarm and Reminder:")
	fmt.Println("ID", alarmID, "is", ac.DistinguishAlarmReminder(alarmID))
	fmt.Println("ID", reminderID, "is", ac.DistinguishAlarmReminder(reminderID))

	// Toggle the alarm off
	ac.ToggleAlarm(alarmID, false)

	// Wait for the alarms and reminders to ring
	time.Sleep(15 * time.Second)

	// Delete an alarm
	ac.DeleteAlarm(alarmID)
}
