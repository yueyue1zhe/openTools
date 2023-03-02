<script lang="ts">
import { computed, defineComponent, h, resolveComponent } from "vue";

export default defineComponent({
  name: "y-icon",
  props: {
    name: {
      type: String,
      required: true,
    },
    size: {
      type: String,
      default: "18px",
    },
    color: {
      type: String,
      default: "",
    },
  },
  setup(props) {
    const iconStyle = computed(() => {
      const { size, color } = props;
      let s = `${size.replace("px", "")}px`;
      return {
        fontSize: s,
        color: color,
      };
    });

    if (props.name.indexOf("el-icon-") === 0) {
      return () =>
        h("el-icon", { class: "icon el-icon", style: iconStyle.value }, [
          h(resolveComponent(props.name)),
        ]);
    } else {
      return () =>
        h("i", { class: [props.name, "icon"], style: iconStyle.value });
    }
  },
});
</script>
<style lang="scss" scoped></style>
