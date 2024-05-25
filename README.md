## GoGrpc-Mis

```mermaid
<!-- RRH -->
flowchart LR
    A[👨‍💼] -->|requests| B{gateway_chan}
    B --> C[buzz_R1_Go]
    B --> D[buzz_R2_Go]
    B --> J[buzz_R3_Go]
```

```mermaid
flowchart LR
    A[👨‍💼] -->|requests| B{gateway}
    B --> C[buzz_Go]
    B --> D[fuzz_Node.js]
```

Benchmark

👉 Fuzz

![fuzz](/assets/fuzz_2024-05-21.png)

👉 Buzz

![buzz](/assets/buzz_2024-05-21.png)

👉 Buzz_RRH

![buzz_rrh](/assets/buzz_rrh_2024-05-25.png)
