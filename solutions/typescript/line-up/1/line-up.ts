export function format(name: string, number: number): string {
  const lastTwo = number%100
  const last = number%10

  let tail = 'th'
  if (lastTwo < 11 || lastTwo >13){
    if (last === 1) tail = 'st'
    else if (last === 2) tail = 'nd'
    else if (last === 3) tail = 'rd'
  }

  return `${name}, you are the ${number}${tail} customer we serve today. Thank you!`
}
