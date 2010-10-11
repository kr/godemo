
                                 Go Statements


                  go f()               ~          setTimeout(f, 0)

                       "At some point in the future, call
                        function f. I don't care exactly
                        when, but the sooner the better."


    * Puts f on a global task queue.

    * Javascript: as long as main keeps running, that's all we'll do. If we
      reutrn all the way up the stack, the top level event loop will pick the
      next thing and start it.

    * Go: can pause main, start f with its own call stack (plus a little other
      context), pause f, resume main, and continue switching back and forth.
      Each context is called a *goroutine*. When the top function in a
      goroutine returns, the goroutine vanishes.

    * This difference makes Go much easier than Javascript to use when mixing
      synchronous and asynchronous code.
