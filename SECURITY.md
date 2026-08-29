# Security

## Secrets

- The repository contains **no API keys, tokens, credentials, or `.env` secrets.**
  `.env` is gitignored; only `.env.example` (variable names with empty values) is
  tracked.
- API keys are supplied through the environment at run time (`SAKSAMA_API_KEY`) and
  are never written to any tracked file.
- Verified: scanning the working tree and full git history for the key prefix
  returns nothing (no `sk-or-...` OpenRouter key in any commit).

## Compromised key notice

During development a free OpenRouter key was pasted into the working session to run
the pipeline. **That key must be considered compromised and should be revoked/rotated**
at <https://openrouter.ai/settings/keys>. It was never committed to the repository.
Do not place any replacement key in the repository — supply it via the environment.

## Data

- All contracts in `data/contracts/` are **synthetic**, written from scratch. No real
  employment contract, and no real personal data, is used. All company and person
  names are fictional.

## Runtime

- The only network calls the engine makes are to the configured LLM API. Evaluation
  (`cmd/eval`) makes **no** network calls and reads only local committed files, so the
  reported numbers can be verified offline with no key.

## Human-in-the-loop

- The output is a triage memo, not a verdict. Every memo states explicitly that it is
  an automatically generated aid, not legal advice, and that the final decision rests
  with the reader.
