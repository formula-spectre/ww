;; http://xahlee.info/emacs/emacs/elisp_syntax_coloring.html

(setq execline-keywords
      (let* (
             (block-management '("foreground" "background" "case" "if" "ifelse" "ifte" "ifthenelse" "backtick" "pipeline" "runblock"))
             (variable-management '("define" "importas" "elglob" "elgetpositionals" "multidefine" "mulstisubstitute"))
             (loops '("forx" "forstdin" "forbacktickx" "loopwhilex")))))

(define-derived-mode execline-mode sh-mode "execline mode"
  "major mode for editing execline scripts"
  (setq font-lock-defaults '((execline-keywords))))
(provide 'execline-mode)
