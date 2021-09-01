import {
  APIGatewayEventRequestContext,
  APIGatewayProxyEvent,
  APIGatewayProxyResult,
} from "aws-lambda";
import { toGLB, toJpeg } from "../lib/converter";
import { fromSVGToString } from "./../lib/file";
export async function handler(
  event: APIGatewayProxyEvent,
  context: APIGatewayEventRequestContext
): Promise<APIGatewayProxyResult> {
  await toJpeg(createSVG(1), "/tmp/sample.jpg");
  toGLB("TEST.gltf", "./result.glb");

  return {
    statusCode: 200,
    body: JSON.stringify({
      id: "",
      method: event.httpMethod,
      query: event.queryStringParameters,
    }),
  };
}

const createSVG = (amount: number) => {
  const svgString = fromSVGToString("sample.svg");
  return svgString.replace("${AMOUNT}", amount.toString());
};
