# CLAUDE.md — Saksama

## Project context

This is a submission for the micro1 Agentic Workflows Hackathon. Hard deadline: 31 August 2026 at 18:00 UTC. Internal checkpoint: 30 August 23:59 UTC. Individual entry. Judges score out of 100 with weights: Agent Solution & Engineering 30, End-to-End Quality 20, Problem & User Value 15, Measured Improvement 15, Reproducibility 15, Hot Take 5. There is a qualification gate: a submission that cannot be run or verified is disqualified before any rubric scoring happens.

The project name is **Saksama** (Indonesian: careful, thorough, meticulous scrutiny). The repository, Go module, and binaries all use the lowercase form `saksama`.

## What we are building

An agent that reviews Indonesian employment contracts (PKWT and PKWTT) for fresh graduates, and reports not only problematic clauses that are present, but also mandatory protections that are missing from the contract.

Intended user: a fresh graduate or entry-level worker in Indonesia who has just received an offer and has two or three days to decide, with no access to legal counsel.

The bottleneck: assessing an employment contract is a judgment call that depends on knowledge of labour law. What traps people is usually not the clause that looks frightening, but the provision that is absent from the document entirely. An unstructured LLM is reactive to text that exists, so it systematically misses absence.

The final output is not a verdict. It is a triage memo containing findings graded by severity, a list of concrete questions to ask HR, and a marker for when to seek legal help. This is not legal advice and the memo must say so explicitly.

## Absolute rules for Claude Code

Do not invent article numbers. Only the fourteen provisions listed under "Legal corpus" may be used. If you believe another provision is relevant, stop and ask; do not add it yourself.

Do not add provisions from Constitutional Court rulings beyond the one listed. Two claims — severance as a mandatory floor, and termination requiring bipartite negotiation — failed verification and are deliberately out of scope.

Do not use real employment contracts sourced from the internet. Every contract in the corpus is written synthetically from scratch.

Do not make any network calls during evaluation other than calls to the LLM API. No fetching legal documents at runtime; all statutes are read from local files.

Do not use an agent orchestration framework (LangChain, LangGraph, or similar). The loop is written by hand in Go.

Do not use the official Anthropic Go SDK. Use `net/http` directly against the Messages API endpoint. The reason is to avoid version uncertainty and protect reproducibility.

Do not use an LLM to grade LLM output. All scoring is deterministic Go code, following the matching rules under "Metrics".

Do not change metric definitions after the baseline has been run.

Do not delete experiments that made results worse. Keep the numbers, record them in the changelog, explain what they taught you. Judges explicitly ask for this.

Do not add dependencies beyond those listed without asking first.

## Stack

Go 1.22 or newer. Standard library only for HTTP and JSON. Permitted: `gopkg.in/yaml.v3` for reading statute and ground truth files. No other dependencies without approval.

The LLM model is configured through the environment variable `SAKSAMA_MODEL`, with the endpoint in `SAKSAMA_API_BASE` and the key in `SAKSAMA_API_KEY`. Do not hardcode a model name anywhere. Do not commit any key; provide a `.env.example` listing variable names with empty values.

## Directory layout

Create exactly this. `cmd/baseline/main.go` for the baseline runner, `cmd/solution/main.go` for the solution runner, `cmd/eval/main.go` for scoring and comparison, `cmd/trajectory/main.go` for trajectory export.

Internal packages: `internal/llm` for the HTTP client and retries, `internal/statutes` for the loader and statute types, `internal/contract` for section splitting, `internal/agent` for the agent stages, `internal/memo` for markdown memo rendering, `internal/scoring` for metrics.

Data: `data/statutes/2026-08.yaml` holding the fourteen provisions, and `data/contracts/c01` through `data/contracts/c12`, each containing `contract.md` and `truth.yaml`.

Outputs: `results/` for stage result JSON, `memos/` for generated memos, `trajectories/` for agent trajectories.

Root documents: `README.md`, `CHANGELOG.md`, `REPRODUCTION.md`, `Makefile`, `Dockerfile`, `.env.example`.

## Legal corpus — fourteen verified provisions

Write to `data/statutes/2026-08.yaml`. Each entry has the fields `id`, `dasar_hukum`, `pasal`, `tier`, `confidence`, `judul`, `ringkasan`, `deteksi`.

Keep field names and all statutory text in Indonesian. These are quotations of Indonesian law and the memo is written for Indonesian readers; translating them would introduce drift and break citation accuracy. Code, comments, and documentation are in English.

`tier` is one of: `batal_demi_hukum`, `sanksi_administratif`, `melanggar_tanpa_sanksi`, `pedoman_kebijakan`.

`confidence` is `A` when verified against primary text, `B` when verified only against consistent secondary sources. This field is mandatory so that the impact of any tier-B error can be isolated from the metrics.

`deteksi` is `ada_klausa` when a violation is detected from the presence of a clause, `tidak_ada_klausa` when detected from the absence of one, or `konteks` when it requires judgment about the nature of the work.

Create these fourteen entries, using the ids exactly as written.

`PP35-12` — PP 35/2021 Pasal 12 — tier `batal_demi_hukum` — confidence A — deteksi `ada_klausa`. PKWT cannot stipulate a probation period. If one is stipulated, it is void by law and the period still counts as length of service.

`UU13-60-1` — UU 13/2003 Pasal 60 ayat 1 — tier `batal_demi_hukum` — confidence B — deteksi `ada_klausa`. PKWTT may stipulate a probation period of at most three months. Beyond that, the excess has no legal force and counts as length of service.

`UU13-60-2` — UU 13/2003 Pasal 60 ayat 2 — tier `batal_demi_hukum` — confidence B — deteksi `ada_klausa`. During probation, the employer may not pay below the applicable minimum wage.

`PP35-4-2` — PP 35/2021 Pasal 4 ayat 2 — tier `melanggar_tanpa_sanksi` — confidence A — deteksi `konteks`. PKWT cannot be used for work that is permanent in nature.

`PP35-8` — PP 35/2021 Pasal 8 — tier `melanggar_tanpa_sanksi` — confidence A — deteksi `ada_klausa`. A term-based PKWT runs at most five years, and the total including any extension may not exceed five years. Critical: this cap applies only to PKWT based on a fixed term under Pasal 5 ayat 1. A PKWT based on completion of specific work is governed by Pasal 9 and has no five-year cap. Reaffirmed in Constitutional Court ruling 168/PUU-XXI/2023.

`PP35-13` — PP 35/2021 Pasal 13 — tier `melanggar_tanpa_sanksi` — confidence A — deteksi `tidak_ada_klausa`. A PKWT must contain at minimum nine items: company name, address, and line of business; worker name, sex, age, and address; position or type of work; place of work; wage amount and payment method; rights and obligations of both parties; start date and term of the PKWT; place and date the PKWT was made; signatures of both parties. Treat these nine as nine separate sub-checks in ground truth.

`PP35-14` — PP 35/2021 Pasal 14 — tier `melanggar_tanpa_sanksi` — confidence A — deteksi `tidak_ada_klausa`. A PKWT must be registered online within three working days of signing.

`PP35-15` — PP 35/2021 Pasal 15 — tier `sanksi_administratif` — confidence A — deteksi `tidak_ada_klausa`. The employer must pay compensation money at the end of the PKWT to any worker with at least one month of continuous service. Does not apply to foreign workers.

`PP35-16` — PP 35/2021 Pasal 16 — tier `melanggar_tanpa_sanksi` — confidence A — deteksi `ada_klausa`. Compensation amount: twelve months of continuous PKWT equals one month of wages; one month or more but under twelve months is prorated as length of service divided by twelve, times one month of wages; over twelve months is likewise prorated. Exception: for micro and small enterprises the amount is set by agreement between the parties.

`PP35-17` — PP 35/2021 Pasal 17 — tier `sanksi_administratif` — confidence A — deteksi `ada_klausa`. If either party ends the employment relationship before the PKWT term expires, the employer still owes compensation calculated on the term actually served.

`PP35-26-31` — PP 35/2021 Pasal 26 and 31 — tier `sanksi_administratif` — confidence A — deteksi `ada_klausa`. Overtime is capped at four hours per day and eighteen hours per week. Overtime pay is 1.5 times the hourly wage for the first hour and 2 times for each subsequent hour. The hourly wage is one one-hundred-seventy-third of the monthly wage.

`PP35-27-5` — PP 35/2021 Pasal 27 ayat 5 — tier `batal_demi_hukum` — confidence A — deteksi `tidak_ada_klausa`. The overtime exemption applies only to defined senior job categories. If those categories are not defined in the employment agreement, company regulations, or the collective agreement, the employer must still pay overtime.

`MK168-79-2b` — Constitutional Court ruling 168/PUU-XXI/2023 on Pasal 79 ayat 2 huruf b of UU 13/2003 as amended by Pasal 81 angka 25 of UU 6/2023 — tier `melanggar_tanpa_sanksi` — confidence A — deteksi `ada_klausa`. The Court held the norm unconstitutional unless read to include the phrase covering two rest days for a five-day working week. Consequently a contract with a five-day week must grant two weekly rest days.

`SE-M5-2025` — Minister of Manpower Circular M/5/HK.04.00/V/2025 dated 20 May 2025 — tier `pedoman_kebijakan` — confidence A — deteksi `ada_klausa`. Employers are prohibited from requiring or holding a worker's diploma or personal documents as security for employment, covering competency certificates, passports, birth certificates, marriage books, and vehicle ownership books. The sole exception is where the diploma or competency certificate was obtained through education or training financed by the employer under a written work agreement, in which case the employer must guarantee the document's safety and compensate the worker if it is damaged or lost. Important note for the agent: a Circular is not binding legislation and carries no sanction, so its enforceability is limited. The memo must state this explicitly rather than calling a diploma-retention clause illegal.

## Contract corpus — twelve synthetic contracts

Every contract is written in Indonesian, in realistic Indonesian employment-contract style, six to twelve pages equivalent, with numbered articles. All company and person names are fictional. Write to `data/contracts/cNN/contract.md`.

`c01` A twelve-month PKWT with a three-month probation clause, plus a clause stating that compensation money is paid only if the contract is extended. Violates PP35-12 and PP35-15.

`c02` A compliant twelve-month PKWT. No violations. Used to measure precision.

`c03` Challenging case. A six-year PKWT that is explicitly based on completion of specific work under Pasal 5 ayat 2, with a clear definition of scope and completion. It appears to breach the five-year cap but is in fact valid because PP35-8 does not apply to this type. Ground truth: no violation on the duration dimension.

`c04` A PKWTT with a six-month probation period. Violates UU13-60-1.

`c05` A PKWT missing four of the nine mandatory items in Pasal 13: place of work, wage payment method, place and date of execution, and company line of business. Violates PP35-13 across four sub-checks.

`c06` A PKWT containing a clause stating the position is not entitled to overtime pay, without defining the senior job category anywhere in the contract. Violates PP35-27-5.

`c07` A PKWT with Monday-to-Friday working hours and one weekly rest day. Violates MK168-79-2b.

`c08` A PKWT requiring surrender of the original diploma as security for the duration of the contract, with no connection to any training. Falls under SE-M5-2025 at the policy-guidance tier.

`c09` A PKWT holding a competency certificate obtained from employer-financed training, with an express safety guarantee and compensation undertaking. This falls within the Circular's exception, so ground truth is: not a violation. Contrast pair for c08.

`c10` A compliant PKWTT with a three-month probation period and wages above minimum wage. No violations.

`c11` Cross-section trap. Article 4 states a twelve-month term, Article 9 states compensation is paid at the end of the term, but Article 14 states that if the worker resigns before the term ends all compensation rights are forfeited. Each reads as reasonable in isolation; combined, Article 14 violates PP35-17. Ground truth must record that detection requires reading across articles.

`c12` A PKWT with a three-month probation period where the probation wage is seventy percent of normal pay and falls below the regional minimum wage. Violates PP35-12 and UU13-60-2.

Final composition: nine contracts with violations, three clean or compliant (c02, c09, c10), one challenging case (c03), one cross-section case (c11).

## Ground truth schema

Write to `data/contracts/cNN/truth.yaml`. Contains `contract_id`, `jenis` valued `PKWT` or `PKWTT`, `skala_usaha` valued `umum` or `mikro_kecil`, and a `findings` array.

Each finding has: `finding_id` unique within the contract, `statute_id` referencing an id from the statute file, `section` holding the contract article number where the problem sits or the literal value `ABSENT` when detection is by absence, `tier` copied from the statute, `deteksi` copied from the statute, `cross_section` valued true or false, and `catatan` holding a one-sentence explanation.

For clean contracts the `findings` array is empty. This is intentional and not a bug.

## Baseline

The baseline is a single LLM call. Send the full contract text in one prompt with an Indonesian-language instruction asking the model to review the employment contract and list the risks it finds. No statute list, no schema, no tools, no retry logic beyond network retries. Output is free-form text.

To make it scoreable, baseline output is normalised through one separate LLM call whose only job is to convert free text into a JSON array with fields `statute_id`, `section`, and `deskripsi`. This normalisation call is given the list of valid statute ids. The normalisation call is not counted as part of the solution and its cost is tracked separately. This matters for a fair comparison and must be explained in the README.

Save to `results/s0_baseline.json`.

## Solution architecture — staged, measured at every stage

Build incrementally and run a full evaluation each time a stage lands. Do not build everything and measure once. Each stage writes its own result file to `results/`.

Stage S1, structured output. Same as baseline but output is forced to JSON with `statute_id` constrained to the fourteen valid ids. Save to `results/s1_structured.json`.

Stage S2, section parser. Split the contract into sections by article numbering in `internal/contract`. Send per section rather than the whole document. Save to `results/s2_sections.json`.

Stage S3, checklist-driven. For every statute with `deteksi` of `ada_klausa` or `konteks`, run one targeted check against the contract. The agent no longer chooses freely what to look at. Save to `results/s3_checklist.json`.

Stage S4, absence pass. For every statute with `deteksi` of `tidak_ada_klausa`, ask explicitly whether the contract contains that provision, requiring a yes or no answer plus the section if present. For PP35-13, run nine separate sub-checks. Save to `results/s4_absence.json`.

Stage S5, citation gate. Every finding with `deteksi` of `ada_klausa` must carry a verbatim quotation from the contract of at most two hundred characters. Verify in Go by substring matching after whitespace normalisation. Findings whose quotation is not found in the contract text are dropped automatically and logged in the trajectory as rejected. Save to `results/s5_gated.json`.

The final solution is S5. `cmd/solution` runs S2 through S5 in sequence.

Every stage must be independently re-runnable via a `-stage` flag.

## Metrics

Define once in `internal/scoring` and do not change after the baseline has run.

Matching is deterministic and implemented in Go. A reported finding matches a ground truth finding when `statute_id` is equal, and additionally, for `ada_klausa` detection, the section number is equal; for `tidak_ada_klausa` detection, matching `statute_id` alone is sufficient. Each ground truth finding may be matched only once. No LLM judgment at this step.

The primary metric is recall: correct findings divided by total ground truth findings across the corpus.

Secondary metrics: precision, meaning correct findings divided by total reported findings. Absence detection rate, meaning recall computed only over findings with `tidak_ada_klausa` detection. Tier accuracy, meaning the proportion of correct findings whose severity tier is also correct. Cross-section recall, meaning recall over findings with `cross_section` true. False positives on clean contracts, meaning the count of findings reported against c02, c09, and c10.

Operational metrics: USD cost per contract, computed from the token usage the API returns, and wall-clock time per contract.

Also compute recall restricted to confidence-A findings only, so that if either confidence-B provision turns out to be wrong its impact can be isolated. Report both numbers.

`cmd/eval` reads two result files, prints a markdown comparison table to stdout, and writes `results/comparison.md`.

## Memo — the final artifact

Render in `internal/memo` to `memos/cNN.md`. Written in Indonesian, because the reader is an Indonesian job candidate. The following structure, in order.

A title with the position and company name taken from the contract. A line giving the assessment date and stating that the assessment is based on law in force as of August 2026, with a note that a new labour protection bill is under deliberation and provisions may change.

An executive summary of at most five sentences, giving the count of findings per tier and a one-sentence overall judgment.

A findings section, grouped by tier in the order: void by law, subject to administrative sanction, in breach without sanction, policy guidance. Each finding carries a short title, the contract article concerned, a verbatim quotation from the contract when detection was clause-based, the full legal basis including article number, two to three sentences on what it means for the worker, and the confidence level A or B.

A section listing provisions not found, enumerating mandatory protections absent from the contract, each with its legal basis.

A section of questions for HR, containing concrete, polite questions that can be sent as-is, one per finding that needs clarification.

A section on when to seek help, appearing only when there is at least one finding at the void-by-law or administrative-sanction tier, advising consultation with the local labour office, a trade union, or a legal aid organisation before signing.

A mandatory closing paragraph stating that this document is an automatically generated triage aid, not legal advice, that it does not replace the judgment of qualified counsel, and that the final decision rests with the reader.

The memo must not read as machine output. No JSON, no technical labels such as `statute_id` or `finding_id`, no emoji, no bullet nesting deeper than one level.

## Reproducibility

`Dockerfile` based on the official Go image, multi-stage, producing a static binary. The container carries no API key; the key is supplied through the environment at run time.

`Makefile` provides targets `build`, `baseline`, `solution`, `eval`, `memos`, `all`, and `clean`. The `eval` target must run without calling the LLM at all when result files already exist in `results/`, so judges can verify the numbers without an API key.

Commit everything in `results/` and `memos/` to the repository. This is what lets judges see the outcome without running anything.

`REPRODUCTION.md` is written for someone starting on a clean machine. It covers the Go and Docker versions used, the exact command sequence from clone through to the comparison table appearing, an explanation of the required environment variables, an estimated runtime and total USD cost for running all six stages, and a sample of the expected output. Also document how to verify results without an API key using the `eval` target.

## Changelog

`CHANGELOG.md` uses a table with columns Stage, What was tried and why, Evidence, and Decision or lesson. One row for the baseline and one row for each stage S1 through S5, plus a final row.

The Evidence column must contain real numbers from `results/`, not qualitative description.

At least one row must cover an experiment that was dropped, with the reasoning. One is already certain: two provisions from the Constitutional Court ruling, on severance as a mandatory floor and on the bipartite negotiation requirement, were deliberately excluded from the corpus because they could not be verified against primary sources within the available time. Record this as a conscious decision, not an oversight.

Close with the main failure mode and the hot take. The intended hot take: an agent that is reactive to text will systematically miss what is absent from that text, and reliability comes from a checklist that forces the agent to ask what is missing, not from a larger model or a cleverer prompt. The baseline versus solution absence detection rate is the evidence for that claim.

## Trajectories

`cmd/trajectory` reads Claude Code session JSONL files from `~/.claude/projects` and converts them into readable markdown in `trajectories/`. Each trajectory must show the agent instructions, the sequence of tool calls with their responses, the feedback that shaped the next step, and any retries or human checkpoints.

Beyond Claude Code trajectories, also capture the application agent's own trajectories. Each run of stages S3 through S5 writes one trajectory file containing the prompt sent, the raw response, the findings that passed the gate, and the findings that were rejected along with the reason.

## README

Open with who the user is and what bottleneck they face, written in prose rather than bullets. Explain why solving it is valuable. Only then cover architecture, usage, and links to `REPRODUCTION.md` and `CHANGELOG.md`.

State explicitly what existed before the competition and what was built during it. For this project, everything was built during it.

Include a short compliance section: data is entirely synthetic, no real personal data is used, no credentials are in the repository, and the solution is designed to keep a human as the final decision maker.

## Order of work

Step 1. Write `data/statutes/2026-08.yaml` with fourteen entries, then `internal/statutes` with unit tests asserting that all ids are unique, all tiers are valid, and every confidence field is populated.

Step 2. Write three contracts first — c01, c02, and c03 — with their ground truth. Stop here and report back for review before writing the remaining nine.

Step 3. Write `internal/llm`, `internal/scoring`, and `cmd/baseline`. Run the baseline over the three contracts. Report the numbers.

Step 4. Complete the corpus to twelve contracts. Re-run the baseline over all twelve. Save as `results/s0_baseline.json`. This is the zero point of the changelog.

Step 5. Build S1 through S5 one at a time. After each stage, run `cmd/eval` against the baseline and record a changelog row. Do not move to the next stage before the current one is measured.

Step 6. Write `internal/memo` and generate all twelve memos. Read one memo end to end yourself and judge whether it reads as human-written. Fix it if it does not.

Step 7. Write the Dockerfile, Makefile, and `REPRODUCTION.md`. Test from a clean container with no cache.

Step 8. Write the README and CHANGELOG. Export trajectories.

## When something is unclear

Stop and ask. Do not fill gaps with assumptions, particularly around article numbers, the content of legal provisions, or metric definitions. Errors in those three areas make every evaluation number meaningless.
