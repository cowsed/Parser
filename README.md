# Parser
An expression parser combined with some math functions and work in progress algebraic simplifications.
At some point in the future the math side of things may be split out into a separate package.

Current features include parsing of +, -, *, /, ^, ln, cos, and sin.

Also can take derivatives* and can do trapezoidal approximations for integrals

\*The derivatives are kind of shoddy currently and are not simplified at all which can lead to problems with readability and NaN appearing when it shouldnt