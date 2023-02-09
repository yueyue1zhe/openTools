export default function toMedia(value: string,prefix:string): string {
    if (!value) return '';
    if (value.includes("http")) return value;
    if (value.indexOf("../") === 0) return value;
    if (value.indexOf('data:') === 0 && value.indexOf('base64') != -1) return value;
    return prefix + value;
}

