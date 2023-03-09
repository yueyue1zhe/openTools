import w7helper from "@/components/yueyue-ui/libs/helper/W7helper";

const w7storage: Storage = {
    length: 0,
    clear: () => {

    },
    getItem: (key: string): string | null => {
        return w7helper.w7.getStorage(key)
    },
    key: (index: number): string | null => {
        return null
    },
    removeItem: (key: string) => {
        w7helper.w7.removeStorage(key);
    },
    setItem: (key: string, value: string) => {
        w7helper.w7.setStorage({key, value});
    }
}

export default w7helper.CanUse() ? w7storage : localStorage