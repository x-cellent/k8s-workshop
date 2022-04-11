Um einen Pod zu skalieren, gibt es ReplicaSets.

Diese sollten aber nicht genutzt werden, da sie normalerweise automatisch gemanaged werden.
Andernfalls könnte ein etwaiges Scaling ohne Wirkung sein.

Erstelle ein Objekt im Namespace `ex5`, welches repliziert UND problemlos gescaled werden kann.
