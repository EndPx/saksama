package scoring

import (
	"fmt"
	"strings"
)

func pct(f float64) string { return fmt.Sprintf("%.1f%%", f*100) }

// Compare renders a markdown table comparing a baseline stage to a solution
// stage. It is what cmd/eval prints to stdout and writes to results/comparison.md.
func Compare(base, sol Metrics) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Comparison: %s vs %s\n\n", base.Stage, sol.Stage)
	fmt.Fprintf(&b, "| Metric | %s | %s |\n", base.Stage, sol.Stage)
	b.WriteString("|--------|------|------|\n")
	row := func(name, a, c string) { fmt.Fprintf(&b, "| %s | %s | %s |\n", name, a, c) }

	row("Recall (primary)", pct(base.Recall), pct(sol.Recall))
	row("Precision", pct(base.Precision), pct(sol.Precision))
	row("Absence detection rate", pct(base.AbsenceRate), pct(sol.AbsenceRate))
	row("Tier accuracy", pct(base.TierAccuracy), pct(sol.TierAccuracy))
	row("Cross-section recall", pct(base.CrossRecall), pct(sol.CrossRecall))
	row("Recall (confidence A only)", pct(base.ConfARecall), pct(sol.ConfARecall))
	row("Correct / reported", fmt.Sprintf("%d / %d", base.Correct, base.ReportedTotal), fmt.Sprintf("%d / %d", sol.Correct, sol.ReportedTotal))
	row("False positives on clean", fmt.Sprintf("%d", base.FalsePositivesClean), fmt.Sprintf("%d", sol.FalsePositivesClean))
	row("Cost per contract (USD)", fmt.Sprintf("$%.4f", base.CostPerContract), fmt.Sprintf("$%.4f", sol.CostPerContract))
	row("Avg time per contract (ms)", fmt.Sprintf("%.0f", base.AvgDurationMs), fmt.Sprintf("%.0f", sol.AvgDurationMs))
	b.WriteString(fmt.Sprintf("\nGround-truth findings in corpus: %d.\n", sol.TruthTotal))
	if base.NormalizationUSD > 0 {
		b.WriteString(fmt.Sprintf("\nBaseline normalisation cost (tracked separately, not part of the solution): $%.4f.\n", base.NormalizationUSD))
	}
	return b.String()
}
