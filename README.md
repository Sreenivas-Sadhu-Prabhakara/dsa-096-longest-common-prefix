# Longest Common Prefix

> Striver SDE Sheet 096 · Day 15 String — carry the first string as the candidate and shrink it against each next string; it can only get shorter. Go, stdlib only, tested.

## Problem

The longest prefix shared by every string in a list (`""` when they share nothing, or the list is empty).

## Approach

The common prefix of n strings is a *fold*: LCP(a, b, c) = LCP(LCP(a, b), c). So hold a candidate — the first string — and against each subsequent string chop the candidate down until it is a prefix (`strings.HasPrefix` is the loop guard). The candidate shrinks monotonically and work is bounded by total characters inspected. Alternatives (vertical character-by-character scanning, sorting and comparing only first/last, binary search on prefix length, a trie) all land in the same complexity class; the fold is the one you can write without thinking, which is its virtue.

**Complexity:** O(total characters) time, O(1) extra space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
