package e2e_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"raja.aiml/ai.explorer/paths"
)

const (
	binaryName = "ai-explorer"
	rootDir    = ".."
)

type TestPaths struct {
	RootDir      string
	BinPath      string
	TemplatePath string
	ConfigPath   string
	PromptOutput string
}

func newTestPaths(promptCategory, topic string) *TestPaths {
	resolver := paths.PathResolver{PromptCategory: promptCategory}
	tmpl, cfg, out := resolver.Derive(topic)

	return &TestPaths{
		RootDir:      rootDir,
		BinPath:      filepath.Join(".build", binaryName),
		TemplatePath: tmpl,
		ConfigPath:   cfg,
		PromptOutput: out,
	}
}

func runCommand(paths *TestPaths, args ...string) ([]byte, error) {
	fullCmd := fmt.Sprintf("%s %s", paths.BinPath, formatArgs(args))
	GinkgoWriter.Println("Running command:")
	GinkgoWriter.Println(fullCmd)

	// Save command to file
	logFilePath := filepath.Join(paths.RootDir, ".build", "executed_prompts.log")
	appendToFile(logFilePath, fullCmd+"\n")

	cmd := exec.Command(paths.BinPath, args...)
	cmd.Dir = paths.RootDir
	cmd.Env = os.Environ()
	return cmd.CombinedOutput()
}

func appendToFile(path, text string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		GinkgoWriter.Printf("Failed to write command to log: %v\n", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(text); err != nil {
		GinkgoWriter.Printf("Failed to write command to log: %v\n", err)
	}
}

func formatArgs(args []string) string {
	quoted := make([]string, len(args))
	for i, arg := range args {
		if needsQuoting(arg) {
			quoted[i] = fmt.Sprintf("%q", arg)
		} else {
			quoted[i] = arg
		}
	}
	return strings.Join(quoted, " ")
}

func needsQuoting(s string) bool {
	return strings.ContainsAny(s, " \t\n\"'")
}

func generatePrompt(paths *TestPaths, category, topic, query string) {
	fullPromptPath := filepath.Join(paths.RootDir, paths.PromptOutput)

	_ = os.Remove(fullPromptPath)
	_ = os.MkdirAll(filepath.Dir(fullPromptPath), 0755)

	args := []string{
		"prompt",
		"--category", category,
		"--topic", topic,
		"--template", paths.TemplatePath,
		"--config", paths.ConfigPath,
		"--output", paths.PromptOutput,
	}

	if query != "" {
		args = append(args, "--query", query)
	}

	output, err := runCommand(paths, args...)
	GinkgoWriter.Printf("CLI output:\n%s\n", string(output))
	Expect(err).ToNot(HaveOccurred(), "Prompt generation command failed")

	Expect(fullPromptPath).To(BeAnExistingFile(),
		fmt.Sprintf("Expected prompt file at: %s", fullPromptPath))
}

var _ = Describe("AI Explorer CLI (E2E)", Ordered, func() {
	DescribeTable("Prompt Generation",
		func(category, topic, query string) {
			paths := newTestPaths(category, topic)
			generatePrompt(paths, category, topic, query)
		},
		Entry("demo category and hello topic", "demo", "hello", ""),
		Entry("topics category and git topic", "topics", "git", ""),
		Entry("topics category and jailbreaking topic", "topics", "jailbreaking", ""),
		Entry("classification/router with query", "classification", "router",
			"Explain Sliding Window Protocol to an CCIE person keep it short"),
	)
})
