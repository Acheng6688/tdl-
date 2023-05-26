package test

import (
	tcmd "github.com/iyear/tdl/cmd"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strconv"
	"testing"
	"time"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}

var (
	cmd         *cobra.Command
	output      string
	testAccount string
)

var _ = BeforeSuite(func() {
	testAccount = strconv.FormatInt(time.Now().Unix(), 10)
})

var _ = BeforeEach(func() {
	cmd = tcmd.New()
	Expect(cmd.PersistentFlags().Set("test", testAccount)).To(Succeed())
	Expect(cmd.PersistentFlags().Set("proxy", "socks5://localhost:1090")).To(Succeed())
})

func exec(cmd *cobra.Command, args []string, success bool) {
	r, w, err := os.Pipe()
	Expect(err).To(Succeed())
	os.Stdout = w

	GinkgoWriter.Printf("args: %s\n", args)
	cmd.SetArgs(args)
	if err = cmd.Execute(); success {
		Expect(err).To(Succeed())
	} else {
		Expect(err).ToNot(Succeed())
	}

	Expect(w.Close()).To(Succeed())

	o, err := io.ReadAll(r)
	Expect(err).To(Succeed())
	output = string(o)
}
