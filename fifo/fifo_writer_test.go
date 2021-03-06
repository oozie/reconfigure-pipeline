package fifo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"os"

	"github.com/pivotal-cf/reconfigure-pipeline/fifo"
)

var _ = Describe("FIFO Writer", func() {
	var (
		writer *fifo.Writer
	)

	BeforeEach(func() {
		writer = fifo.NewWriter()
	})

	It("writes to a fifo and returns the path", func() {
		contents := "My Special Contents"

		path, err := writer.Write(contents)
		Expect(err).NotTo(HaveOccurred())

		defer os.Remove(path)

		readContents, err := ioutil.ReadFile(path)
		Expect(err).NotTo(HaveOccurred())

		Expect(readContents).To(Equal([]byte(contents)))
	})
})
