import * as grpc from "@grpc/grpc-js";

import { handler } from "../handler";
import { FuzzServiceService } from "../genproto/fuzz";

export function runGrpcServer(add: string): void {
	const server = new grpc.Server();

	server.addService(FuzzServiceService, handler);

	server.bindAsync(
		`0.0.0.0:${add}`,
		grpc.ServerCredentials.createInsecure(),
		(err: Error | null, port: number) => {
			if (err) {
				console.error(`Server error: ${err.message}`);
			} else {
				console.log(`Grpc server running on port: ${port}`);
				server.start();
			}
		}
	);
}
