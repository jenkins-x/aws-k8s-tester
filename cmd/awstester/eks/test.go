package eks

import (
	"fmt"
	"os"

	"github.com/aws/awstester/eksconfig"
	"github.com/aws/awstester/eksdeployer"
	"github.com/aws/awstester/internal/eks"

	"github.com/spf13/cobra"
)

func newTest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test <subcommand>",
		Short: "Test commands",
	}
	cmd.AddCommand(
		newTestALB(),
	)
	return cmd
}

func newTestALB() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alb <subcommand>",
		Short: "Test ALB commands",
	}
	cmd.AddCommand(
		newTestALBCorrectness(),
		newTestALBQPS(),
		newTestALBMetrics(),
	)
	return cmd
}

func newTestALBCorrectness() *cobra.Command {
	return &cobra.Command{
		Use:   "correctness",
		Short: "Runs ALB correctness test",
		Run:   testALBCorrectness,
	}
}

func testALBCorrectness(cmd *cobra.Command, args []string) {
	if path == "" {
		fmt.Fprintln(os.Stderr, "'--path' flag is not specified")
		os.Exit(1)
	}

	cfg, err := eksconfig.Load(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load configuration %q (%v)\n", path, err)
		os.Exit(1)
	}
	var dp eksdeployer.Interface
	dp, err = eks.NewEKSDeployer(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create EKS deployer %v\n", err)
		os.Exit(1)
	}

	if err = dp.TestALBCorrectness(); err != nil {
		fmt.Fprintf(os.Stderr, "failed correctness test %v\n", err)
		os.Exit(1)
	}
}

func newTestALBQPS() *cobra.Command {
	return &cobra.Command{
		Use:   "qps",
		Short: "Runs ALB QPS test",
		Run:   testALBQPS,
	}
}

func testALBQPS(cmd *cobra.Command, args []string) {
	if path == "" {
		fmt.Fprintln(os.Stderr, "'--path' flag is not specified")
		os.Exit(1)
	}

	cfg, err := eksconfig.Load(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load configuration %q (%v)\n", path, err)
		os.Exit(1)
	}
	var dp eksdeployer.Interface
	dp, err = eks.NewEKSDeployer(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create EKS deployer %v\n", err)
		os.Exit(1)
	}

	if err = dp.TestALBQPS(); err != nil {
		fmt.Fprintf(os.Stderr, "failed scalability QPS test %v\n", err)
		os.Exit(1)
	}
}

func newTestALBMetrics() *cobra.Command {
	return &cobra.Command{
		Use:   "metrics",
		Short: "Runs ALB metrics test",
		Run:   testALBMetrics,
	}
}

func testALBMetrics(cmd *cobra.Command, args []string) {
	if path == "" {
		fmt.Fprintln(os.Stderr, "'--path' flag is not specified")
		os.Exit(1)
	}

	cfg, err := eksconfig.Load(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load configuration %q (%v)\n", path, err)
		os.Exit(1)
	}
	var dp eksdeployer.Interface
	dp, err = eks.NewEKSDeployer(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create EKS deployer %v\n", err)
		os.Exit(1)
	}

	if err = dp.TestALBMetrics(); err != nil {
		fmt.Fprintf(os.Stderr, "failed metrics test %v\n", err)
		os.Exit(1)
	}
}