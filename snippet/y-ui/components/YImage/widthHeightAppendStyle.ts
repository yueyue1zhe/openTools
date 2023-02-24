export default function widthHeightAppendStyle(width: number | string, height: number | string) {
    let outStyle = ";";
    let useWidth = uni.$y.addUnit(width);
    let useHeight = uni.$y.addUnit(height);
    outStyle += `height:${useHeight ? useHeight : useWidth};`;
    outStyle += `width:${useWidth ? useWidth : useHeight};`;
    return outStyle;
}