// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>

package kata
import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    . "codewarrior/kata"
)

func dotest(x, n int, expected []int) {
    Expect(CountBy(x, n)).To(Equal(expected), "With x = %d, n= %d", x, n)
}

var _ = Describe("Tests", func() {
    It("Sample tests", func() {
        dotest(1, 5, []int{1, 2, 3, 4, 5})
        dotest(2, 5, []int{2, 4, 6, 8, 10})
        dotest(3, 5, []int{3, 6, 9, 12, 15})
        dotest(50, 5, []int{50, 100, 150, 200, 250})
        dotest(100, 5, []int{100, 200, 300, 400, 500})
    })
})