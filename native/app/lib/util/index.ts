/**
 * 全角の英数字、ピリオド、カンマを含んだ文字列を半角の英数字、ピリオド、カンマに変換する関数
 * @param  {String} 変換前の文字列
 * @return {String} 変換後の文字列
 */
export function fullWidth2halfWidth(str: string): string {
  return str.replace(/[Ａ-Ｚａ-ｚ０-９，．]/g, (s) => {
    return String.fromCharCode(s.charCodeAt(0) - 0xFEE0);
  });
}
