;;; execline-mode.el --- major mode for editing execline scripts. -*- coding: utf-8; lexical-bindings: t -*-


;;; Commentary:
;;; small major mode for editing scripts written in the execline language, made by skarnet.
;;; https://skarnet.org/software/execline

;; Code
(setq execline-keywords
      (let* (
             (process-state-control '("execline-cd" "execline-umask" "posix-cd" "posix-umask" "emptyenv" "envfile" "export" "fdblock" "fdclose" "fdmove" "fdreserve" "fdswap" "redirfd" "piperw" "heredoc" "wait" "getcwd" "getpid" "exec" "tryexec" "exit" "trap" "withstdinas"))
             (basic-block-management '("foreground" "background" "case" "if" "ifelse" "ifte" "ifthenelse" "backtick" "pipeline" "runblock"))
             (variable-management '("define" "importas" "elglob" "elgetpositional" "multidefine" "multisubstitute"))
             (loops '("forx" "forstdin" "forbacktickx" "loopwhilex"))
             (positional-parameters '("elgetopt" "shift" "dollarat"))
             (misc '("eltest" "homeof")))

             ;; generate regexes for each category
             (block-management-regexp (regexp-opt block-management 'words))
             (variable-management-regexp (regexp-opt block-management 'words))
             (loops-regexp (regexp-opt loops 'words)))
        `(
          (,block-management-regexp . 'font-lock-keyword-face)
          (,variable-management-regexp . 'font-lock-keyword-face)
          (,loops-regexp . 'font-lock-keyword-face)))


(define-derived-mode execline-mode sh-mode "execline mode"
  "Major mode for editing execline scripts."
  (setq font-lock-defaults '((execline-keywords))))
(provide 'execline-mode)

;;; execline-mode.el ends here
