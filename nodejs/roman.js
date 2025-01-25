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
        let quotient = 0;
        let remainder = 0;

        let subres = [];
        
        if (divident < 10) {
            res.push(mapa[divident])
        } else {
            let counter = 0;
            let str = `${numbers[i]}`           
            for (let j = 10;counter < str.length;j *=10){
                quotient = parseInt(divident / j);
                remainder = parseInt(divident % j);
                
                if (mapa[remainder] !== undefined) {
                    subres.unshift(mapa[remainder])
                    divident = divident - remainder
                } else {
                    if (remainder !== 0) {
                        
                        let notInMapa = [];
                        let transformRemainer = remainder;
                        for (let k=0;k<remainder;k++){
                            if (mapa[transformRemainer] == undefined){
                                notInMapa.unshift(mapa[j / 10])
                            } else {
                                notInMapa.unshift(mapa[transformRemainer])
                                break
                            }
                            transformRemainer = transformRemainer - j / 10;
                        }
                        
                        subres.unshift(notInMapa.join(""))
                        
                        divident = divident - remainder
                        
                    } else {
                        if (mapa[divident] !== undefined) {
                            subres.unshift(mapa[divident])
                        } else {
                            let notInMapa = [];
                            let transformDivident = divident;
                            for (let k=0;k<divident;k++){
                                if (mapa[transformDivident] == undefined){
                                    notInMapa.unshift(mapa[j])
                                } else {
                                    notInMapa.unshift(mapa[transformDivident])
                                    break
                                }
                                transformDivident = transformDivident - j;
                            }
                            subres.unshift(notInMapa.join(""))
                        }
                    }
                }
                counter++
            }

            res.push(subres.join(""))
        }
        
    } 
    console.log(res)
    return res
}

function run() {
    romanizer([80]);
}

run()