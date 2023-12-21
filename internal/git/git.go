package git

import (
	"fmt"
	"os"
	"os/exec"
)

type Git struct {
}

func New() *Git {

	g := &Git{}

	return g
}

func (g *Git) IsInstalled() bool {

	_, err := exec.LookPath("git")

	return err == nil

}

func (g *Git) Commit(message string) error {

	gitCmd := exec.Command("git", "add", ".")
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr
	if err := gitCmd.Run(); err != nil {
		return fmt.Errorf("git ci: %w", err)
	}

	gitCmd = exec.Command("git", "commit", "--amend", "-m", message)
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr
	if err := gitCmd.Run(); err != nil {
		return fmt.Errorf("git ci: %w", err)
	}

	return nil
}

func (g *Git) Add(workDir string, repository string, ref string, message string) error {

	gitCmd := exec.Command("git", "subtree", "add", "--prefix", workDir, "--message", message, repository, ref, "--squash")
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr
	if err := gitCmd.Run(); err != nil {
		return fmt.Errorf("git subtree add: %w", err)
	}

	return nil

}

func (g *Git) UpStream(workDir string, repository string, ref string) error {

	gitCmd := exec.Command("git", "subtree", "push", "--prefix", workDir, repository, ref)
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr
	if err := gitCmd.Run(); err != nil {
		return fmt.Errorf("git subtree push: %w", err)
	}

	return nil
}

func (g *Git) Update(workDir string, repository string, currRef string, message string, targetRef string) error {

	gitCmd := exec.Command("git", "subtree", "pull", "--prefix", workDir, "--message", message, repository, targetRef, "--squash")
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr
	if err := gitCmd.Run(); err != nil {
		return fmt.Errorf("git subtree pull: %w", err)
	}

	return nil

}

func (g *Git) Remove(workDir string, message string) error {

	gitCmd := exec.Command("git", "rm", "-rf", workDir)
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr
	if err := gitCmd.Run(); err != nil {
		return fmt.Errorf("git rm: %w", err)
	}

	gitCmd = exec.Command("git", "commit", "--message", message)
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr
	if err := gitCmd.Run(); err != nil {
		return fmt.Errorf("git commit: %w", err)
	}

	return nil
}

func (g *Git) Log() ([]byte, error) {

	cmd := exec.Command("git", "log", "--grep=^git-vendor-name:", "--pretty=format:START %H%n%s%n%n%b%nEND%n", "HEAD")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("git log: %w", err)
	}

	return output, nil

}
