<script lang="ts">
	import Button from '../../../lib/components/button/Button.svelte';
	import Input from '../../../lib/components/input/Input.svelte';
	import { createContext } from '../../../lib/utils/createContext';
	import z from 'zod';
	import { toaster } from '$lib/stores/toaster';
	import type { AgentContext, CommentFormContext } from './types.ts';
	import type { iComment } from '$lib/types/comment';
	import validateSchema from '$lib/utils/validateSchema';

	export let parentComment: iComment | undefined = undefined;
	$: isAReply = !!parentComment;
	let text: string;
	let errors = new Map<string, string>();
	let loading = false;
	const formSchema = z.object({
		text: z
			.string()
			.min(1, 'Comment must have at least 1 digit')
			.max(500, 'Comment must not have more than 500 digits')
	});

	const agentContext = createContext<AgentContext>("agent").get()
	const agent = agentContext.agent

	const onSubmit = async (e: SubmitEvent) => {
		loading = true;
		try {
			e.preventDefault()
			if(!!!agent.uuid) {
				toaster.add({
				title: 'Error',
				description: 'Something went wrong',
				type: 'error'
			});
			} 
			const comment = { agent_id: agent.uuid, text, parent_id: isAReply ? parentComment?.id : undefined  };
			const form = { text }
			const result = validateSchema(formSchema, form);
			if (result.type === 'error') {
				errors = result.errors;
				return;
			}
			errors = new Map();
			
		} catch (error) {
			toaster.add({
				title: 'Error',
				description: 'Could not create an account, try again later',
				type: 'error'
			});
		}
		loading = false;
	};

	const contextKey = `comment-form-${parentComment ? parentComment.id : Math.random()}`;
	const context = createContext<CommentFormContext>(contextKey);
	const contextValue = { formSchema };
	context.set(contextValue);
</script>

<form
	on:submit={onSubmit}
	enctype="multipart/form-data"
	class="flex flex-col gap-1"
	data-testid="form"
>
	<Input
		{errors}
		formContext={context}
		label={isAReply ? `Reply to @mushz` : 'Create a comment'}
		bind:value={text}
		name="comment"
		containerClassName="max-w-[500px]"
	/>
	<Button {loading} intent={isAReply ?  "ghost" : "primary"} type="submit"  id="submit-btn">Submit</Button>
</form>
