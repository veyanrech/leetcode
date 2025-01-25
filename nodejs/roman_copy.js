function romanizer(numbers) {
    // Write your code here
    const mapa = {
        1:"I",          
        2:"II",           
        3:"III",          
        4:"IV",          
        5:"V",          
        6:"VI",          
        7:"VII",          
        8:"VIII",         
        9:"IX",
        10:"X",
        40:"XL",
        50:"L",
        90:"XC",
        100:"C",
        400:"CD",
        500:"D",
        900:"CM",
        1000:"M"
    }
    
    const res = [];

    for (let i = 0; i < numbers.length; i++) {
        
        let divident = numbers[i];
        let j = 10;
        let subres = [];
        while (divident > 0) {
            let remainder = parseInt(divident % j);

            console.log(remainder)
            if (remainder == 0) {
                j *= 10;
                continue;
            }

            if (mapa[remainder] !== undefined) {
                subres.unshift(mapa[remainder])
            } else {

                let transformReminder = remainder;

                while (transformReminder > 0) {
                    if (mapa[transformReminder] == undefined) {
                        subres.unshift(mapa[j / 10])
                    } else {
                        subres.unshift(mapa[transformReminder])
                        break
                    }
                    transformReminder = transformReminder - j / 10;
                    console.log(subres)
                }
            }

            divident = divident - remainder

            j *= 10;

            console.log(subres, j, divident)

        };

        res.push(subres.join(''));
    }

    console.log(res)

    return res;
}

function run() {
    romanizer([80, 75, 161]);
}

run()