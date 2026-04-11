type StaticImage = {
  src: string;
  width: number;
  height: number;
  blurDataURL?: string;
};

declare module '*.svg' {
  const content: StaticImage;
  export default content;
}

declare module '*.png' {
  const content: StaticImage;
  export default content;
}

declare module '*.jpg' {
  const content: StaticImage;
  export default content;
}

declare module '*.jpeg' {
  const content: StaticImage;
  export default content;
}

declare module '*.gif' {
  const content: StaticImage;
  export default content;
}

declare module '*.webp' {
  const content: StaticImage;
  export default content;
}
