namespace YOptForm {
    interface OptsItemType {
        name: string;
        type: number;
        label?: string;
        placeholder?: string;
        required?: boolean;

        uploadImageOption?: UploadImageOption;
    }

    interface UploadImageOption {
        toMediaFunc: ToMediaFunc
        actionFunc: ()=>Promise<AttachResult>
    }
}