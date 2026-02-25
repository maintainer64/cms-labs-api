// @ts-ignore
import parseCurl from 'parse-curl';

export const curlParsedData = (raw?: string) => {
  const curlParsed = parseCurl(raw);
  if (!curlParsed) {
    throw new Error('Curl not parsed data');
  }
  return curlParsed;
};

export const curlParsedDataValidator = (raw?: string) => {
  try {
    curlParsedData(raw);
    return null;
  } catch (e) {
    return e?.toString() || 'Curl not parsed';
  }
};
