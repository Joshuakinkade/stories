import SignUpForm from '@/components/signupForm';

export default async function SignUpPage({ params }: { params: Promise<{ code: string }> }) {
  const { code } = await params;
  console.log(code);

  return (
    <div>
      <div className="w-96 mx-auto">
      <SignUpForm code={code} />
      </div>
    </div>
  )
}
