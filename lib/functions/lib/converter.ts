import svgToImg from "@lildiary/svg-to-img";
import { ConvertGltfToGLB } from "gltf-import-export";
import chronium = require("chrome-aws-lambda");

import path = require("path");

export const toGLB = (gltbFileName: string, outputFileName: string): void => {
  ConvertGltfToGLB(pathFromAssetName(gltbFileName), outputFileName);
};

const pathFromAssetName = (filename: string) => {
  return path.join(__dirname, "..", "voucher", "assets", filename);
};

export const toJpeg = async (svg: string, dist: string) => {
  let browser = null;
  await svgToImg
    .from(svg)
    .toJpeg({
      path: dist,
    })
    .catch((e) => {
      console.log(e);
    })
    .then((t) => {
      console.log(t);
    })
    .finally(() => console.log("finished"));
};
