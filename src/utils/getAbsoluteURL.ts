const getAbsoluteURL = (path: string): string =>
  `${process.env.NEXT_PUBLIC_URL}${path}`

export default getAbsoluteURL
