<script>
  import dayjs from "dayjs";
  export let from;
  export let to;

  $: a = dayjs(from);
  $: b = dayjs(to);
  $: relative = a.to(b, true);
  $: duration = dayjs.duration(b.diff(a));
  $: human = duration.as("minute") + " minute(s)";
  $: computer = duration.toISOString();
</script>

<time datetime={computer} title={human}>{relative}</time>
