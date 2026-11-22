export const detectMimeType = (base64: string) => {
  if (base64.startsWith("/9j/")) {
    return "image/jpeg";
  } else if (base64.startsWith("iVBOR")) {
    return "image/png";
  } else if (base64.startsWith("R0lGOD")) {
    return "image/gif";
  } else if (base64.startsWith("UklGR")) {
    return "image/webp";
  }
  return "image/jpeg";
};

export const toImageSrc = (image: string) => {
  if (!image) {
    return;
  }
  const value = image.trim();
  if (value.startsWith("data:")) {
    return value;
  } else {
    return `data:${detectMimeType(value)};base64,${value}`;
  }
};
