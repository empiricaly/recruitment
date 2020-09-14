<script>
  import dayjs from "dayjs";
  export let from;
  export let to;

  $: a = dayjs(from);
  $: b = dayjs(to);
  $: relative = a.to(b, true);
  $: duration = dayjs.duration(b.diff(a));
  $: minutes = Math.round(duration.as("minute"));
  $: human = minutes + ` minute${minutes == 1 ? "" : "s"}`;
  $: output = minutes < 120 ? human : relative;
  $: computer = duration.toISOString();
</script>

<time datetime={computer} title={human}>{output}</time>
