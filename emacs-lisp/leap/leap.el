;;; leap.el --- Leap exercise (exercism)  -*- lexical-binding: t; -*-

;;; Commentary:

(defun leap-year-p (year)
  (or (= (mod year 400) 0)
      (and (= (mod year 4) 0)
           (/= (mod year 100) 0))))

(provide 'leap-year-p)
;;; leap.el ends here
