import * as grpc from "@grpc/grpc-js";

import { PingRequest, PongResponse } from "../genproto/buzz";

function ping(
	call: grpc.ServerUnaryCall<PingRequest, PongResponse>,
	callback: grpc.sendUnaryData<PongResponse>
) {
	// const reqParams = call.request;
	callback(null, { str: "Pong" });
}

export const handler = {
	ping,
};
