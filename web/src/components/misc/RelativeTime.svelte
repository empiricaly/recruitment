<script>
  import dayjs from "dayjs";
  export let time;

  $: now = dayjs();
  $: relative = now.to(time);
  $: dtime = dayjs(time);
  $: human = dtime.calendar();
  $: computer = dtime.toISOString();
</script>

<time datetime={computer} title={relative}>{human}</time>
