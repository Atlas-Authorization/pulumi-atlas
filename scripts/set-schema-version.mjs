#!/usr/bin/env node
// Set the top-level "version" of the generated Pulumi schema.json.
//
// tfgen omits "version", but the Pulumi registry needs it. This edits the text
// (not a JSON re-serialize) so tfgen's formatting is kept byte-for-byte.
//
//   node scripts/set-schema-version.mjs <schema.json> <version>
import { readFileSync, writeFileSync } from 'node:fs';

const [file, version] = process.argv.slice(2);
if (!file || !/^\d+\.\d+\.\d+(-[0-9A-Za-z.-]+)?$/.test(version ?? '')) {
  console.error('usage: set-schema-version.mjs <schema.json> <semver>');
  process.exit(2);
}
const text = readFileSync(file, 'utf8');
const doc = JSON.parse(text); // fail loudly on a broken schema
const top = /^ {2}"version": "[^"]*"/m; // two-space indent = a top-level key
let out;
if ('version' in doc && top.test(text)) {
  out = text.replace(top, `  "version": ${JSON.stringify(version)}`);
} else {
  const end = text.lastIndexOf('}');
  if (end < 0) throw new Error(`${file}: no closing brace`);
  out = `${text.slice(0, end).trimEnd()},\n  "version": ${JSON.stringify(version)}\n${text.slice(end)}`;
}
if (JSON.parse(out).version !== version) throw new Error('version not set');
writeFileSync(file, out);
console.log(`${file}: version ${version}`);
