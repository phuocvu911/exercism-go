//
// This is only a SKELETON file for the 'Line Up' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const format = (name, num) => {
  const lastTwo = num%100
  const last = num%10

  let tail = "th"
  if (lastTwo < 11 || lastTwo >13){
    switch (last){
      case 1: 
        tail = "st"
        break
      case 2:
        tail = "nd"
        break
      case 3:
        tail = "rd"
        break
    }
  }
  return `${name}, you are the ${num}${tail} customer we serve today. Thank you!`
};
