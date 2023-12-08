import { readdir, writeFile, readFile } from "node:fs/promises";

const DIR = process.cwd() + "/src/assets/images/";
const NAME_PREFIX = "IMG_";
const OLD_FILE_DATA = [];
const NEW_ADDED = [];
let CURRENT_WRITTEN = "";

async function getAllImages() {
  const files_from_dir = await readdir(DIR);
  console.log(`Generating At Directory = ${DIR}`);
  const images = files_from_dir.filter((file) =>
    /.(png|jgp|svg|webp|mp4|gif)/g.test(file)
  );
  return images;
}

async function newAdded() {
  const newAdded = CURRENT_WRITTEN.split(/\n/gi).filter(
    (d) => !OLD_FILE_DATA.includes(d)
  );

  newAdded.forEach((d) => NEW_ADDED.push(d));

  if (NEW_ADDED.length === 0) return console.log(`😒 No New file Added`);

  NEW_ADDED.forEach((e) => console.log("✅ added ", e));
}

async function makingExports() {
  try {
    const allImages = await getAllImages();

    //Storing Previous File Data;
    (await readFile(DIR + "index.ts"))
      .toString("utf8")
      .split(/\n/gi)
      .forEach((old) => OLD_FILE_DATA.push(old));

    let text = "";

    allImages.forEach((img_path, index) => {
      let img_name = img_path
        .replace(/.(png|jgp|svg|webp|mp4|gif)/g, "")
        .replace(/-/g, "_")
        .replace(/\s/g, "")
        .toUpperCase();
      text += `export { default as ${NAME_PREFIX}${img_name} } from "./${img_path}";
`;

      if (index === allImages.length - 1) {
        CURRENT_WRITTEN = text;
        writeFile(`${DIR}/index.ts`, text)
          .then(async (value) => {
            newAdded();
          })
          .catch((e) => console.error(e));
      }
    });
  } catch (e) {
    console.log(`⛔ got Error ${e.message}`);
  }
}

makingExports();
