
const writeWords =(options: WriteWordsOptions) => {

    options.ctx.setFontSize(options.fontSize);//设置字体大小

    var allRow = Math.ceil(options.ctx.measureText(options.word).width / options.maxWidth);//实际总共能分多少行

    var count = allRow >= options.maxLine ? options.maxLine : allRow;//实际能分多少行与设置的最大显示行数比，谁小就用谁做循环次数

    var endPos = 0;//当前字符串的截断点

    for (var j = 0; j < count; j++) {

        var nowStr = options.word.slice(endPos);//当前剩余的字符串

        var rowWid = 0;//每一行当前宽度

        if (options.ctx.measureText(nowStr).width > options.maxWidth) {//如果当前的字符串宽度大于最大宽度，然后开始截取

            for (var m = 0; m < nowStr.length; m++) {

                rowWid += options.ctx.measureText(nowStr[m]).width;//当前字符串总宽度

                if (rowWid > options.maxWidth) {

                    if (j === options.maxLine - 1) { //如果是最后一行

                        options.ctx.fillText(nowStr.slice(0, m - 1) + '...', options.x, options.y + (j + 1) * options.fontSize); //(j+1)*fontSize这是每一行的高度

                    } else {

                        options.ctx.fillText(nowStr.slice(0, m), options.x, options.y + (j + 1) * options.fontSize);

                    }

                    endPos += m;//下次截断点

                    break;

                }

            }

        } else {//如果当前的字符串宽度小于最大宽度就直接输出
            let appendX = (options.maxWidth - options.ctx.measureText(nowStr).width) / 2
            options.ctx.fillText(nowStr.slice(0), options.x + appendX, options.y + (j + 1) * options.fontSize);

        }

    }

}

/**该方法用来绘制一个有填充色的圆角矩形
 *@param cxt:canvas的上下文环境
 *@param x:左上角x轴坐标
 *@param y:左上角y轴坐标
 *@param width:矩形的宽度
 *@param height:矩形的高度
 *@param radius:圆的半径
 *@param fillColor:填充颜色
 **/
function fillRoundRect(cxt:UniNamespace.CanvasContext, x:number, y:number, width:number, height:number, radius:number, /*optional*/ fillColor:string) {
    //圆的直径必然要小于矩形的宽高
    if (2 * radius > width || 2 * radius > height) { return false; }

    cxt.save();
    cxt.translate(x, y);
    //绘制圆角矩形的各个边
    drawRoundRectPath(cxt, width, height, radius);
    cxt.fillStyle = fillColor || "#000"; //若是给定了值就用给定的值否则给予默认值
    cxt.fill();
    cxt.restore();
}
/**该方法用来绘制圆角矩形
 *@param cxt:canvas的上下文环境
 *@param x:左上角x轴坐标
 *@param y:左上角y轴坐标
 *@param width:矩形的宽度
 *@param height:矩形的高度
 *@param radius:圆的半径
 *@param lineWidth:线条粗细
 *@param strokeColor:线条颜色
 **/
function strokeRoundRect(cxt:UniNamespace.CanvasContext, x:number, y:number, width:number, height:number, radius:number, /*optional*/ lineWidth:number, /*optional*/ strokeColor:string) {
    //圆的直径必然要小于矩形的宽高
    if (2 * radius > width || 2 * radius > height) { return false; }

    cxt.save();
    cxt.translate(x, y);
    //绘制圆角矩形的各个边
    drawRoundRectPath(cxt, width, height, radius);
    cxt.lineWidth = lineWidth || 2; //若是给定了值就用给定的值否则给予默认值2
    cxt.strokeStyle = strokeColor || "#000";
    cxt.stroke();
    cxt.restore();
}

function drawRoundRectPath(cxt:UniNamespace.CanvasContext, width:number, height:number, radius:number) {
    cxt.beginPath();
    //从右下角顺时针绘制，弧度从0到1/2PI
    cxt.arc(width - radius, height - radius, radius, 0, Math.PI / 2);

    //矩形下边线
    cxt.lineTo(radius, height);

    //左下角圆弧，弧度从1/2PI到PI
    cxt.arc(radius, height - radius, radius, Math.PI / 2, Math.PI);

    //矩形左边线
    cxt.lineTo(0, radius);

    //左上角圆弧，弧度从PI到3/2PI
    cxt.arc(radius, radius, radius, Math.PI, Math.PI * 3 / 2);

    //上边线
    cxt.lineTo(width - radius, 0);

    //右上角圆弧
    cxt.arc(width - radius, radius, radius, Math.PI * 3 / 2, Math.PI * 2);

    //右边线
    cxt.lineTo(width, height - radius);
    cxt.closePath();
}


export default {
    writeWords,
    fillRoundRect
}