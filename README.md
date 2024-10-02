# CVmaker

## Steps to Install Latex in Ubuntu

## Latex Install
- `sudo apt-get update`
- `sudo apt-get install texlive-full`
- {Recommended}: `sudo apt-get install texlive-latex-base texlive-latex-recommended texlive-latex-extra`

## Debug Latex Install
- `sudo fmtutil-sys --all`
- `sudo apt --fix-broken install`
- `sudo apt-get remove --purge texlive*`
- `sudo apt-get autoremove`
- `sudo apt-get install texlive-latex-base texlive-latex-recommended texlive-latex-extra`

## Verify
- `which pdflatex` Expected Output: `/usr/bin/pdflatex`