import express, { Request, Response } from "express";
import cors from "cors";
// import promClient from "prom-client";

import { runGrpcServer } from "./grpc";

const app = express();
// const register = new promClient.Registry();

// register.setDefaultLabels({
// 	app: "monitoring-article",
// });

app.use(cors());
const port = 9000;

// app.get("/metrics", async (req: Request, res: Response) => {
// 	res.setHeader("Content-Type", register.contentType);
// 	res.send(await register.metrics());
// });

app
	.listen({ port }, async () => {
		console.log(`Server listening ${port}`);
	})
	.on("error", (error) => {
		throw new Error(error.message);
	});

runGrpcServer("5003");
