import 'server-only';
import { cookies } from 'next/headers';
import { getIronSession } from 'iron-session';

const SESSION_PASSWORD = process.env.SESSION_PASSWORD || '';

export interface Session {
  userEmail: string;
  userId: string;
  authToken: string;
}

export async function loadSession(): Promise<Session> {
  return await getIronSession<Session>(await cookies(), { password: SESSION_PASSWORD, cookieName: "auth" });
}

export async function startSession(session: Session) {
  const ttl = 24 * 60 * 60;
  const ironSession = await getIronSession<Session>(await cookies(), {
    password: SESSION_PASSWORD, cookieName: "auth", ttl, cookieOptions: {
      sameSite: "strict"
    }
  });
  ironSession.authToken = session.authToken;
  ironSession.userId = session.userId;
  ironSession.userEmail = session.userEmail;
  await ironSession.save();
}

export async function destroySession(): Promise<null> {
  const session = await getIronSession<Session>(await cookies(), { password: SESSION_PASSWORD, cookieName: "auth" });
  session.destroy();
  return null;
}
