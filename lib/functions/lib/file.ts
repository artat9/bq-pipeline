import path = require("path");
import * as fs from "fs";

export const fromSVGToString = (filename: string): string => {
  const buffer = fs.readFileSync(path.join(__dirname, "assets", filename));
  return buffer.toString();
};
