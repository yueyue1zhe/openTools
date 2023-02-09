
export default function splitArray<T>(arr: T[], len:number){
    let a_len = arr.length;
    let result = [];
    for(let i = 0 ; i < a_len ; i += len){
        result.push( arr.slice( i, i + len ));
    }
    return result;
}