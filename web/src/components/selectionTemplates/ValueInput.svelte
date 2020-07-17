<script context="module">
  function getType(value) {
    if (value.int) {
      return "int";
    } else if (value.float) {
      return "float";
    } else if (value.string) {
      return "string";
    } else if (value.boolean) {
      return "boolean";
    }
  }

  function isFloat(value) {
    let splittedValues = value.split(".");

    if (splittedValues.length !== 2) {
      return false;
    }

    for (let i = 0; i < splittedValues.length; i++) {
      if (!Number.isInteger(+splittedValues[i])) {
        return false;
      }
    }

    return true;
  }

  function isInteger(value) {
    return Number.isInteger(+value);
  }

  function isBoolean(value) {
    return value === "true" || value === "false";
  }

  function setValue(value) {
    value = value.trim();
    if (value === "") {
      return;
    }

    if (isFloat(value)) {
      return { float: parseFloat(value) };
    } else if (isInteger(value)) {
      return { int: parseInt(value) };
    } else if (isBoolean(value)) {
      return { boolean: value === "true" };
    } else {
      return { string: value };
    }
  }
</script>

<script>
  import Input from "../form/Input.svelte";

  export let criteria = {};

  function mapValues() {
    return criteria.values.map(v => v[getType(v)]).join(", ");
  }

  let value = criteria.values && mapValues();

  function handleValueOnKeyUp(event) {
    const rawValue = event.target.value;
    let values = rawValue.split(",");
    values = values.map(v => setValue(v));
    values = values.filter(c => c !== undefined && c !== null);
    criteria.values = values;
  }

  function handleValueOnBlur(event) {
    criteria.values = criteria.values.filter(
      c => c !== undefined && c !== null
    );
    value = mapValues();
  }
</script>

{#if criteria.values}
  <Input
    bind:value
    placeholder="value"
    onBlur={handleValueOnBlur}
    onKeyUp={handleValueOnKeyUp} />
{:else}
  <div class="w-full bg-gray-100 border border-gray-300 px-3 py-2 rounded-md" />
{/if}
