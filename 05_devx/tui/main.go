package tui

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	currentTime    string
	nextMeeting    string
	tasks          []string
	testStatus     string
	lintWarnings   string
	packageUpdates []string
	systemStats    string
	networkInfo    string
	quickCommands  []string
	vaultSecrets   string
	tip            string
}

func InitialModel() model {
	return model{
		currentTime:    time.Now().Format("03:04 PM"),
		nextMeeting:    "11:00 AM - Team Standup",
		tasks:          []string{"[ ] Refactor login API", "[ ] Write unit tests for /auth", "[ ] Update README with usage"},
		testStatus:     "✅ All tests passed",
		lintWarnings:   "2 warnings",
		packageUpdates: []string{"- github.com/gin-gonic/gin (v1.9.1 → v1.10)", "- golang.org/x/tools (v0.12 → v0.13)"},
		systemStats:    "CPU: 32%   Mem: 58%   Disk: 74% used",
		networkInfo:    "IP: 192.168.1.5   Ping google.com: 21ms",
		quickCommands:  []string{"1. Run Tests", "2. Git Pull", "3. Open Notes.md"},
		vaultSecrets:   "2 Secrets Stored",
		tip:            `Type "devdash ai 'how to write middleware in Go'"`,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	var builder strings.Builder

	// Header Section
	builder.WriteString(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205")).Render("🚀 DevDash v0.1.0") + "\n")
	builder.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("════════════════════════════════════════════════════════════════════════════\n"))

	// Project Info
	builder.WriteString(fmt.Sprintf("📁 Project: backend-api  | 🌱 Git: feature/login | 🔄 Pull needed\n"))

	// Time and Meeting
	builder.WriteString(fmt.Sprintf("🕒 Time: %s        | 📅 Next Meeting: %s\n", m.currentTime, m.nextMeeting))
	builder.WriteString("════════════════════════════════════════════════════════════════════════════\n")

	// Tasks
	builder.WriteString("📋 Tasks Today:\n")
	for _, task := range m.tasks {
		builder.WriteString(fmt.Sprintf("  %s\n", task))
	}
	builder.WriteString("════════════════════════════════════════════════════════════════════════════\n")

	// Code Status and Package Updates
	builder.WriteString("🧪 Code Status:              📦 Package Updates:\n")
	builder.WriteString(fmt.Sprintf("  %s\n", m.testStatus))
	builder.WriteString(fmt.Sprintf("  🔍 Lint: %s\n", m.lintWarnings))
	for _, update := range m.packageUpdates {
		builder.WriteString(fmt.Sprintf("  %s\n", update))
	}
	builder.WriteString("════════════════════════════════════════════════════════════════════════════\n")

	// System Stats and Network Info
	builder.WriteString("🖥️  System Stats:                   🌐 Network Info:\n")
	builder.WriteString(fmt.Sprintf("  %s\n", m.systemStats))
	builder.WriteString(fmt.Sprintf("  %s\n", m.networkInfo))
	builder.WriteString("════════════════════════════════════════════════════════════════════════════\n")

	// Quick Commands
	builder.WriteString("🧠 Quick Commands:\n")
	for _, cmd := range m.quickCommands {
		builder.WriteString(fmt.Sprintf("  %s\n", cmd))
	}
	builder.WriteString("════════════════════════════════════════════════════════════════════════════\n")

	// Vault and Tips
	builder.WriteString(fmt.Sprintf("🔒 Vault: %s | Type `devdash vault` to manage\n", m.vaultSecrets))
	builder.WriteString(fmt.Sprintf("🤖 Tip: %s\n", m.tip))
	builder.WriteString("════════════════════════════════════════════════════════════════════════════\n")

	return builder.String()
}

func Start() {
	p := tea.NewProgram(InitialModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error running TUI:", err)
		os.Exit(1)
	}
}
